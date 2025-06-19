package main

import "fmt"

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

type Number interface { // ประกาศ interface ที่ระบุว่าเป็นจำนวน
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func min[T Number](a, b T) T { // Generics ควรใช้เมื่อมีการสร้าง func ที่เหมือนกัน แล้วต้องการให้มันอยู่ใน func เดียว
	if a < b {
		return a
	}
	return b
}

func main() {
	var a, b = 1, 2
	fmt.Println("minInt a,b: ", minInt(a, b))

	c, d := 1.1, 2.2
	fmt.Println("minFloat64 c,d: ", minFloat64(c, d))

	fmt.Println("min a,b: ", min(a, b))
	fmt.Println("min c,d: ", min(c, d))
}
