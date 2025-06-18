package main

import (
	"errors"
	"fmt"
)

//error เป็น intrerface นึง ซึ่งจะถูกนิยามไว้แล้ว
/*type error interface {
	Error() string
}*/

type MyError struct{}

func (e MyError) Error() string {
	return "Cannot divide by zero"
}

var MyErr = errors.New("my custom error message.")

func Divide(a, b float64) (float64, error) { // ต้อง return 2 ค่า
	if b == 0 {
		// err := MyError{}
		// err := fmt.Errorf("can't divide '%f' by zero", a) //สามารถกำหนด string เองได้
		return 0, MyErr // ต้อง return 2 ค่า
	}
	r := a / b
	return r, nil // ต้อง return 2 ค่า // zero value ของ error คือ nil
}

func main() {
	r, err := Divide(1, 0)
	if err != nil {
		fmt.Println("handler err : ", err)
		return // หยุดการทำงาน (ไม่รันคำสั่งบรรทัดถัดไป)
	}
	fmt.Println("r : ", r)
}
