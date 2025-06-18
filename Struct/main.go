package main

import (
	"fmt"
)

type course struct { // class ในภาษาอื่นๆ
	name       string
	instructor string
	price      float64
}

func main() {
	c := course{
		name:       "Go",
		instructor: "John",
		price:      100,
	} // การประกาศ Struct แบบที่ 1

	var c1 = course{"Rust", "Rick", 450} // การประกาศ Struct แบบที่ 2

	var c2 = course{
		name: "C++",
	} // การประกาศ Struct แบบที่ 3

	var c3 = course{} // การประกาศ Struct แบบที่ 4 zero value

	fmt.Println("name : ", c.name) // เข้าถึงตัวแปรโดยใช้ dot (.)
	fmt.Println("instructor : ", c.instructor)
	fmt.Println("price : ", c.price)

	c.instructor = "Nong" // เปลี่ยนค่าใน Struct
	fmt.Println("new instructor : ", c.instructor)

	fmt.Printf("c : %#v\n", c)
	fmt.Printf("c1 : %#v\n", c1)
	fmt.Printf("c2 : %#v\n", c2)
	fmt.Printf("c3 : %#v\n", c3) //zero value
}
