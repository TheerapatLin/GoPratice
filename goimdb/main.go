package main

// libraly "ECHO"
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// go mod init github.com/TheerapatLin/goimdb
	// go mod tidy // go get
	_ "github.com/proullon/ramsql/driver"
)

type Movie struct {
	ID          int64   `json:"id"`
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

func getAllMoviesHandler(c echo.Context) error {
	mvs := []Movie{}
	y := c.QueryParam("year")

	if y == "" {
		rows, err := db.Query(`SELECT id, imdbID, title, year, rating, isSuperHero
		FROM goimdb`)
		if err != nil {
			log.Fatal("query error", err)
		}
		defer rows.Close()

		for rows.Next() {
			var m Movie
			if err := rows.Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero); err != nil {
				return c.JSON(http.StatusInternalServerError, "scan:"+err.Error())
			}
			mvs = append(mvs, m)
		}

		if err := rows.Err(); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, mvs)
	}

	year, err := strconv.Atoi(y)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	rows, err := db.Query(`SELECT id, imdbID, title, year, rating, isSuperHero
	FROM goimdb
	WHERE year = ?`, year)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var m Movie
		if err := rows.Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		mvs = append(mvs, m)
	}

	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, mvs)
}

func getMovieByIDHandler(c echo.Context) error {
	imdbID := c.Param("imdbID")
	fmt.Println("id : ", imdbID)

	row := db.QueryRow(`select * from goimdb where imdbID=?`, imdbID)
	m := Movie{}
	err := row.Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, m)
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "not found",
		})
	default:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func createMovieHandler(c echo.Context) error {
	m := &Movie{}
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// เปิด statement
	stmt, err := db.Prepare(` 
	INSERT INTO goimdb (imdbID,title,year,rating,isSuperHero) 
	VALUES (?,?,?,?,?)
	`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close() //  ควรปิดหลังเปิด statement เสมอ

	// b := fmt.Sprintf("%v", m.IsSuperHero) // fmt.SprintF() convert boolean to string
	r, err := stmt.Exec(m.ImdbID, m.Title, m.Year, m.Rating, m.IsSuperHero)
	switch {
	case err == nil:
		id, _ := r.LastInsertId() // ดึง id ที่ insert ล่าสุด
		m.ID = id
		return c.JSON(http.StatusCreated, m)
	case err.Error() == "UNIQUE constraint violation":
		return c.JSON(http.StatusConflict, "movie already exists")
	default:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// return c.JSON(http.StatusCreated, m) // 201 = http.StatusCreated
}

var db *sql.DB

func conn() {
	var err error
	db, err = sql.Open("ramsql", "goimdb")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	conn() // connect to DB

	createTb := `
	CREATE TABLE IF NOT EXISTS goimdb (
	id INT AUTO_INCREMENT,
	imdbID TEXT NOT NULL UNIQUE,
	title TEXT NOT NULL ,
	year INT NOT NULL ,
	rating FLOAT NOT NULL ,
	isSuperHero BOOLEAN NOT NULL ,
	PRIMARY KEY (id)
	);`

	if _, err := db.Exec(createTb); err != nil { // create table
		log.Fatal("create table error : ", err)
	}

	// create new echo instance
	e := echo.New()            // สร้าง struct echo
	e.Use(middleware.Logger()) // middleware สำหรับดู log request

	e.GET("/movies", getAllMoviesHandler)
	e.GET("/movies/:imdbID", getMovieByIDHandler) // /:id คือ path params

	e.POST("/movies", createMovieHandler)

	PORT := "2565"

	log.Println("starting... port : ", PORT)

	log.Fatal(e.Start(":" + PORT)) // start server

}
