package main

import (
	"fmt"
	"strings"
)

// maps คือ data structure โครงสร้างข้อมูลประเภทหนี่ง key & value // key ห้ามซ้ำกัน

func main() {
	var ss []string = []string{}            // ประกาศ slice และ set zero value
	var m map[string]int = map[string]int{} // ประกาศ map และ set zero value // key ของ map คือ string, value คือค่าที่เก็บอยู่ใน map เป็น int (ชนิดข้อมูลเป็นแบบไหนก็ได้)
	fmt.Printf("ss : %#v\n", ss)
	fmt.Printf("m : %#v\n", m)

	m["Answer"] = 42
	fmt.Printf("new m : %#v\n", m)

	v := m["Answer"] // read map
	fmt.Println("v : ", v)

	m["Nong"] = 15
	fmt.Printf("new m : %v\n", m)

	delete(m, "Answer") // delete ค่าใน map
	fmt.Printf("new m : %v\n", m)

	fmt.Println("len(m) : ", len(m))             // นับจำนวนค่าใน map
	fmt.Println("find 'Answer' : ", m["Answer"]) // result คือ 0

	fmt.Printf("new m : %#v\n", m)

	n, ok := m["Answer"] // ตรวจสอบว่ามีค่าใน map หรือไม่
	if ok {
		fmt.Println("found : ", n)
	} else {
		fmt.Println("not found")
	}

	s := "If it looks like a duck swims like a duck and quacks like a duck then it probably is a duck "
	w := WordCount(s)
	fmt.Printf("w : %#v", w)

	// fmt.Println("w")
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	fmt.Printf("words :  %#v\n", words)
	r := map[string]int{}
	for _, w := range words { // _ คือ index
		r[w] = r[w] + 1
		fmt.Println("r[w] : ", r[w])
		fmt.Println("w : ", w)
	}
	return r
}
