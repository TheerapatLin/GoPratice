package main

import (
	"fmt"
	"strconv" // package สำหรับการแปลง string
)

func main() {
	var i int = 256
	fmt.Printf("type i : %T, val : %d\n", i, i)

	var f float64 = float64(i) // แปลงค่า
	fmt.Printf("type f : %T, val : %f\n", f, f)

	v := "72"
	s, err := strconv.Atoi(v) // func นี้จะ reeturn 2 ค่าเสมอ
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("s : %#v\n", s)

	ii := 10
	n := strconv.Itoa(ii) // แปลงค่า int เป็น string
	fmt.Printf("n : %#v\n", n)
}
