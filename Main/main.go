package main

import (
	"fmt"
	"math"
)

func main() {

	//variable
	fmt.Println(`Variable 
	`)
	var msg string = "msg" // กำหนด string หรือไม่ก็ได้
	fmt.Println("msg : ", msg)
	fmt.Printf("msg format : %s\n", msg) // print format, %s : string

	msg1 := "msg1" // ประกาศตัวแปรแบบสั้นๆ แต่ไม่สามารถกำหนดไว้นอก func ได้
	fmt.Println("msg1 : ", msg1)

	msg = "msg edit" + " again" // สามารถเปลี่ยนค่าได้ แต่ยังคงเป็น string เท่านั้น
	fmt.Println("msg edit : ", msg)

	/*
		คอมเม้นหลายบรรทัด
	*/

	var age int = 55
	fmt.Println("age : ", age)
	fmt.Printf("age format : %d\n", age) // print format, %d : integer

	var price float64 = 22.512
	fmt.Println("price : ", price)
	fmt.Printf("price format : %.3f\n", price) // print format, %f : float , ทศนิยม 3 ตำแหน่ง

	var check bool = true
	fmt.Println("check : ", check)
	fmt.Printf("check format : %t\n", check)

	cost, lifepoint := 100, false
	fmt.Println("cost : ", cost, "\n", "lifepoint : ", lifepoint)

	movie := ` 	
movie : {
	title : tenet
	year : 2020
	rating : 7.5
	category : sci-fi
	hero : false
}
	` // raw string
	fmt.Println(movie)

	var r rune = 'L' // ASCII int32, rune ใช้ '' ได้
	fmt.Println("r : ", r)
	fmt.Printf("r format : %c\n", r) // print format, %c : character

	fmt.Printf("msg : %#v\n", msg)                   // %#v ใช้ print ได้ทุกค่า
	fmt.Printf("type : %T -- age : %#v\n", age, age) // %T ระบุ type ของตัวแปร

	//zero value
	var msg2 string
	var age1 int
	var price1 float64
	var check1 bool
	var r1 rune
	fmt.Printf("msg2 : %s\n", msg2)
	fmt.Printf("msg2 q : %q\n", msg2) //%q สามารถแสดง ""
	fmt.Printf("age1 : %d\n", age1)
	fmt.Printf("price1 : %.2f\n", price1)
	fmt.Printf("check1 : %t\n", check1)
	fmt.Printf("r1 : %d\n", r1)

	//if-else
	fmt.Println("\nIf-else")
	num := 35

	if num == 34 {
		fmt.Println("\nnum == 34")
	} else {
		fmt.Println("\nnum != 34")
	}

	if num == 34 && (num > 30 || num <= 39) {
		fmt.Println("Yes!! it's thirty four.")
	}

	if num%2 == 0 {
		fmt.Println("num is even")
	} else if num == 35 {
		fmt.Println("whatt")
	} else {
		fmt.Println("num is odd")
	}

	//short statement
	lim := 225.0
	if v := math.Pow(10, 2); v < lim { // v ไม่สามารถถูกเรียกนอก if ได้ เนื่องจาก {}
		fmt.Println("x power n is : ", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
	}

	//switch case
	fmt.Println("\nSwitch-case")

	switch today := "Sunday"; today { //ไม่จำเป็นต้องใส่ break
	case "Monday", "Tuesday", "Wednesday":
		fmt.Println("\ntoday is weekdays")
	case "Sunday":
		fmt.Println("\ntoday is Sunday")
		fallthrough //ทำงาน case ต่อไปด้วย แค่ case ถัดไปเท่านั้น
	case "Saturday":
		fmt.Println("today is Saturday")
	default:
		fmt.Println("\ntoday is whattttttt")
	}

	switch {
	case num < 0:
		fmt.Println("num is negative")
	case num == 0:
		fmt.Println("num is zero")
	case num > 0:
		fmt.Println("num is positive")
	default:
		fmt.Println("num is whattttttt")
	}

	//function
	add(12, 5)
	a, b := add(12, 89)
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

	c, d := swap("hello", "world")
	fmt.Println(c, d)

	split(4) // naked return

	fmt.Println("funcVariable : ", add1(78, 5))
	fmt.Printf("%T\n", add1)

	r2 := compute(add1)
	fmt.Println("r2 : ", r2)

	x1 := compute(hypot) // function ที่มี parameter เป็น function
	fmt.Println("x1 : ", x1)

	//higher order function
	fmt.Println("\nHigher order function")

	inc, curr := adder()
	v2 := inc()
	fmt.Println(v2)
	v2 = inc()
	fmt.Println(v2)

	y1 := curr()
	fmt.Println(y1)

	//Array
	fmt.Println("\nArray")

	var skills [3]string = [3]string{"js", "go", "python"} //เก็บข้อมูลชนิดเดียวกัน
	fmt.Printf("%#v\n", skills)

	s := skills[1]
	fmt.Printf("s = %#v\n", s)

	l := len(skills)
	fmt.Printf("l = %#v\n", l)

	skills[1] = "golang"
	fmt.Printf("skills = %#v\n", skills)
	show(skills)

	var xs [3]string
	show(xs)

	genres := [...]string{"action", "horror", "comedy"} // ไม่ต้องกำหนดขนาด array แล้ว
	fmt.Printf("%#v\n",genres)

}

// function
func add(x float64, y float64) (float64, string) { //parameter ,ต้องกำหนดประเภทตัวแปรที่ return ด้วย ถ้า return 2 ค่า ต้องใส่วงเล้บ
	fmt.Println("add", x, y)
	v := x + y
	return v, "hello"
}

func swap(x string, y string) (string, string) {
	return y, x
}

// naked return
func split(sum int) (x, y int) { //อย่าใช้กับ func ยาวๆ
	x = sum * 4 / 9
	y = sum - x
	return
}

var add1 func(float64, float64) float64 = func(x, y float64) float64 {
	return x + y
}

func compute(fn func(float64, float64) float64) float64 {
	fmt.Println("Compute!!")
	v1 := fn(42, 13)
	return v1
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// higher order function
func adder() (func() int, func() int) {
	sum := 0
	return func() int {
			sum = sum + 1
			return sum
		}, func() int {
			return sum
		}
}

func show(sk [3]string) {
	fmt.Println("show : ", sk)
	sk[2] = "ruby"
}
