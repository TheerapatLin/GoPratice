package main

import "fmt"

const Pi = 3.14

type day int

func main() {
	const world = "世界" // ตัวแปร const คือตัวแปรที่ไม่สามารถเปลี่ยนค่าใหม่ได้
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const (
		_          = iota // iota สามารถใส่ค่าตัวเลขตามลำดับของค่าถัดไป โดยเพิ่มขึ้นทีละ 1
		Sunday day = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday :", Sunday)
	fmt.Println("Monday :", Monday)
	fmt.Println("Tuesday :", Tuesday)
	fmt.Println("Wednesday :", Wednesday)
	fmt.Println("Thursday :", Thursday)
	fmt.Println("Friday :", Friday)
	fmt.Println("Saturday :", Saturday)

	skills("js", "go", "python")
}

func skills(xs ...string) {
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

	for _, s := range xs {
		fmt.Println("skills : ", s)
	}
}
