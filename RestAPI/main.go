package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies []Movie

func movieHandler(w http.ResponseWriter, req *http.Request) { // http.ResponseWriter คือ interface // *http.Request คือ struct
	method := req.Method // req.Method ,GET คือ http method
	if method == "POST" {

		body, err := ioutil.ReadAll(req.Body) // อ่าน body ของ req
		if err != nil {
			fmt.Fprintf(w, "error : %v", err) // นำ string ที่กำหนดส่งไปกับ response
			return
		}
		fmt.Println("Body : ", string(body))

		t := Movie{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // ระบุที่ header ว่า bad request // ต้องแปลง json เปน struct ก่อน
			fmt.Fprintf(w, "error : %s", err)    // นำ string ที่กำหนดส่งไปกับ response
			return
		}

		movies = append(movies, t)

		fmt.Fprintf(w, "Hello %s created movies", "POST")
		return
	}

	b, err := json.Marshal(movies) // แปลง struct เป็น json
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error : %s", err) // นำ string ที่กำหนดส่งไปกับ response
		return
	}

	w.Header().Set("content-type", "application/json; charset=utf-8") // set header ให้ content-type เป็น json

	w.Write(b) // นำ struct ที่แปลงเป็น json ส่งไปกับ response (แบบย่อ)
	// fmt.Fprintf(w, "%s", string(b)) // นำ struct ที่แปลงเป็น json ส่งไปกับ response

}

func main() {

	/*
		GET เรียก
		POST สร้าง
		PUT อัพเดต แก้ไข
		DELETE ลบ
	*/

	// ถ้ามี req เข้ามาที่ path นั้นๆ จะเรียกใช้ func movieHandler
	http.HandleFunc("/movies", movieHandler)          // http.HandleFunc รับพารามิเตอร์ 2 ตัวคือ path และ handler (func ที่ไว้จัดการ req,res)
	err := http.ListenAndServe("localhost:2565", nil) // รัน server ที่ port:2565
	log.Fatal(err)
}
