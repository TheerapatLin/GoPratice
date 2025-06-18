package main

import (
	"encoding/json"
	"fmt"
)

type Course struct { // ต้องเปลี่ยน key เป็นตัวพิมพ์ใหญ่
	Name       string `json:"name"`
	Level      string `json:"level"`
	Views      int    `json:"views"`
	Instructor string `json:"instructor"`
	FullPrice  int    `json:"fullprice"` // ทำให้สามารถใช้ตัวพิมพ์เล็กได้ // ต้องเขียนติดกันเท่านั้น
}

type movie struct {
	Title       string `json:"title"`
	Year        int
	Rating      float32
	Genres      []string
	IsSuperHero bool
}

func main() {
	c := Course{ // struct // ต้องเปลี่ยน key เป็นตัวพิมพ์ใหญ่ ไม่่งั้น json จะไม่สามารถเข้าถึงได้
		Name:       "Basic GO",
		Level:      "normal",
		Views:      7840,
		Instructor: "Anuchito",
		FullPrice:  1200,
	}

	b, err := json.Marshal(c) // Marshal รับค่าเป็น any , return เป็น byte array
	fmt.Printf("type : %T\n", b)
	fmt.Printf("byte : %v\n", b)
	fmt.Printf("byte to string : %v\n", string(b))
	fmt.Printf("string : %s\n", b)
	fmt.Println(err)

	data := []byte(`{"name":"Advanced Python","level":"advanced","views":40,"instructor":"Anuchito","fullprice":11200}`)
	var c1 Course
	err1 := json.Unmarshal(data, &c1)
	fmt.Printf("%#v\n", c1)
	fmt.Println(err1)

	body := `[
	{
		"imdbID":"tt4154796",
		"title":"Batman",
		"year":1989,
		"rating":7.6,
		"genres":["Action","Adventure"],
		"isSuperHero":true
	},
	{
		"imdbID":"tt4154756",
		"title":"Joker",
		"year":2018,
		"rating":8.4,
		"genres":["Action","Drama"],
		"isSuperHero":false
	}
	]`

	ms := []movie{}

	err2 := json.Unmarshal([]byte(body), &ms)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("%#v\n", ms)
}
