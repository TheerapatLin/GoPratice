package main

import (
	"fmt"
)

func show(val int) {
	fmt.Println(val)
}

func showNew(val any) {
	// i, ok := val.(int)
	// if ok {
	// 	i = i * 2
	// 	fmt.Println(i)
	// } else {
	// 	fmt.Println("not int")
	// }
	i, _ := val.(int) // "_" คือตัวแปร checking
	i = i + 2
	fmt.Println(i)

	switch v := val.(type) {
	case int:
		i := v + 3
		fmt.Println("int:", i)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("ไม่รู้จักประเภทข้อมูล")
	}
}

type promotioner interface { // interface สามารถใช้ในการรับค่าได้ และจำเป็นต้องมีเพียง Method ที่กำหนดไว้เท่านั้น หาก Struct นั้นๆ มี Method มากกว่าก็ไม่เป็นไร
	discount() int // ตัวแปรใดๆก็ตามที่เข้ามาใน func นั้นๆด้วย interface สามารถเรียก method แค่ของ interface เท่านั้น แต่สามารถถูกเียกใช้ Method อื่นใน func main ได้
} // การตั้งชื่อ interface แนะนำให้ตั้งชื่อลงท้ายด้วย er เช่น promotioner

type infoer interface {
	info()
}

type presenter interface {
	promotioner
	infoer
}

func summary(val presenter) {
	fmt.Println("discount price : ", val.discount())
	val.info()
}

func sale(val promotioner) {
	fmt.Printf("sale : %#v\n", val.discount())
}

func showInfo(val infoer) {
	val.info()
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info : ", c)
}

func main() {
	var v interface{} // empty interface ตัวแปรที่สามารถเก็บ type ไหนก็ได้
	v = 36
	fmt.Printf("v : %T %#v\n", v, v)
	v = "hello"
	fmt.Printf("v : %T %#v\n", v, v)

	var v1 any // (เฉพาะเวอชั่น 1.8 ขึ้นไป) empty interface ตัวแปรที่สามารถเก็บ type ไหนก็ได้
	v1 = 36
	fmt.Printf("v1 : %T %#v\n", v1, v1)

	show(v1.(int)) // เปลี่ยน type จาก any เป็น int
	showNew(v)
	v2 := course{}
	sale(v2)
	v2.info()
	showInfo(v2)
	summary(v2)
}
