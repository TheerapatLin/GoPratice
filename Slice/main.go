package main

import "fmt"

func main() {
	skills := []string{"js", "go", "python"} //slice สามารถยืดหดขนาดได้
	fmt.Printf("(without #) %T => %v\n", skills, skills)
	fmt.Printf("(with #) %T => %#v\n", skills, skills) //ระบุ type ของ skills

	ss := append(skills, "java", "c++") //append สามารถเพิ่มข้อมูลต่อท้ายได้
	fmt.Printf("length: %d -- val:%#v\n", len(skills), skills)
	fmt.Printf("length: %d -- val:%#v\n", len(ss), ss)

	s1 := skills[0:2] //slice ตัดข้อมูล
	show(s1)
	/*
		[0:3) => {0,1,2} เอา 0 ไม่เอา 3
		[0:2) =>  {0,1} เอา 0 ไม่เอา 2
		[1:3) =>  {1,2} เอา 1 ไม่เอา 3
	*/
	l := len(ss)
	show(ss[0:l]) //แสดง index แรกถึง index สุดท้าย
	show(ss[0:])  //แสดง index แรกถึง index สุดท้าย
	show(ss[:l])  //แสดง index แรกถึง index สุดท้าย
	show(ss[:])   //แสดง index แรกถึง index สุดท้าย

	var ss1 []int
	fmt.Printf("%#v\n", ss1) //zero value คือ nil

	ss2 := make([]int, 3) //make เป็นการจองพื้นที่ใน slice
	fmt.Printf("%#v\n", ss2)

	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}
	var votes []float64
	votes = append(xs, ys...) //รวม 2 slice
	fmt.Printf("%#v\n", votes)

	s3 := skills[0:2]
	showw("s3", s3)

	s4 := skills[1:3]
	showw("s4", s4)

	s3[1] = "c#" // หากมีการเปลี่นนค่าใน index ใน slice ของ slice หรือ Array เดิม จะเปลี่่ยน slice หรือ Array เดิมนั้นๆด้วย
	showw("s3", s3)
	showw("s4", s4)
	fmt.Printf("skills: %#v\n", skills)

	s4[0] = "ruby"
	showw("s3", s3)
	showw("s4", s4)
	fmt.Printf("skills: %#v\n", skills)

}

func show(sk []string) {
	fmt.Printf("show : %#v\n ", sk)
}

func showw(tag string, sk []string) {
	l := len(sk)  // มีกี่ index
	c := cap(sk)	// สามารถเก็บได้กี่ index
	fmt.Printf("%s : length: %d capacity: %d-- show:%v\n", tag, l, c, sk)
}
