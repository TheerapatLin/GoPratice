package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/proullon/ramsql/driver" // _ หน้าสุดมีหน้าที่ที่จะไม่ลบตัว import นี้ออกหลังเซฟ
)

func main() {
	db, err := sql.Open("ramsql", "goimdb") // sql.Open() คือ การจองความจำเพื่อเก็บข้อมูลใน db โดยตั้งชื่อ db เป็น goimdb // db คือ instance ที่ไว้คุยกับ database
	if err != nil {
		log.Fatal("error ramsql : ", err)
		return
	}

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

	_, err = db.Exec(createTb) // db.Exec() คือ รันคำสั่ง db
	if err != nil {
		log.Fatal("error createTb : ", err)
		return
	}

	fmt.Println("table created.")

	insert := `
	insert into goimdb(imdbID,title,year,rating,isSuperHero)
	values(?,?,?,?,?);
	`

	stmt, err := db.Prepare(insert) // เตรียม statement ให้เรียกใช้งาน
	if err != nil {
		log.Fatal("prepare staement error : ", err)
	}
	r, err := stmt.Exec("tt123456", "Batman", 2016, 8.6, true)
	if err != nil {
		log.Fatal("insert error : ", err)
	}
	l, err := r.LastInsertId() // id ล่าสุดที่ถูกเก็บ
	fmt.Println("lastInsert : ", l, "\terr l : ", err)
	ef, err := r.RowsAffected() // มีผลกระทบกี่บรรทัด
	fmt.Println("RowAffected : ", ef, "\terr ef : ", err)

	rows, err := db.Query("select id,imdbID,title,year,rating,isSuperHero from goimdb") // Query คือ การดึงข้อมูลมาเก็บไว้ใน rows // ตัวแปร rows เป็น struct
	if err != nil {
		log.Fatal(" query error : ", err)
		return
	}

	for rows.Next() { // rows.Next() คือการเช็คว่าใน rows มีของมั้ย return เป็น boolean
		var id int
		var imdbID, title string
		var year int
		var rating float32
		var isSuperHero bool

		err := rows.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero) // Scan คือ การเก็บข้อมูลที่ได้จาก query ไว้ในตัวแปร
		if err != nil {
			log.Fatal("scan error : ", err)
			return
		}
		fmt.Println("row : ", id, imdbID, title, year, rating, isSuperHero)
	}

	stmt2, err := db.Prepare(`
	UPDATE goimdb
	SET rating = ?
	WHERE imdbID = "tt123456"
	;`)
	if err != nil {
		log.Fatal("prepare update error : ", err)
	}
	_, err = stmt2.Exec(9.9)
	fmt.Println("stmt2 exec error : ", err)
	if err != nil {
		log.Fatal("update error : ", err)
	}
	fmt.Println(stmt2)

	rowx := db.QueryRow("select id, imdbID, title, year, rating, isSuperHero from goimdb where imdbID=?", "tt123456")
	var id int
	var imdbID, title string
	var year int
	var rating float32
	var isSuperHero bool
	err = rowx.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
	if err != nil {
		log.Fatal("scan one rowx error : ", err)
		return
	}

	fmt.Println("one rowx : ", id, imdbID, title, year, rating, isSuperHero)

}
