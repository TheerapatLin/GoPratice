package main

import "fmt"

func changePrice(price int) {
	price = price - 599
	fmt.Println("change :", price, &price)
}

type course struct {
	name, instructor string
	price            int
}

func discount(c *course) int { // หากต้องการเปลี่ยนค่าตัวแปรใน struct ตรงๆต้องใช้ pointer เพื่อระบุที่อยู่ของตัวแปร (c *course)
	c.price = c.price - 599
	fmt.Println("discount func : ", c.price)
	return c.price
} // เหมือนกับการใช้ method

type movie struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m *movie) addVote(rating float64) []float64 { // (m *movie)
	m.votes = append(m.votes, rating)
	fmt.Printf("m.votes : %#v\n", m.votes)
	// var addr *[]float64 = &m.votes
	// *addr = append(*addr, rating)
	// fmt.Printf("addr : %#v\n", addr)
	// fmt.Printf("*addr : %#v\n", *addr)
	fmt.Printf("%#v\n", m)
	return m.votes
}

func main() {

	var a = 10
	var addra *int = &a                                      // เก็บที่อยู่ของค่า a int
	fmt.Printf("a : %v\nPointer's a (addra):%v\n", a, addra) // &a (& ตามด้วยตัวแปร) คือการแสดงที่อยู่ใน Memory ของตัวแปรนั้นๆ (Memory address)

	price := 9999
	addr := &price // addr มี type เป็น *int (Pointer)
	fmt.Printf("Type's addr : %T,%v\n", addr, addr)
	*addr = 9400                                       // write เขียนค่าใหม่ในที่อยู่นั้นๆ
	fmt.Printf("price edited : %v, %v\n", price, addr) // price เปลี่ยนค่าด้วย เนื่องจาก addr คือที่อยู่ของ price
	v := *addr                                         // read อ่านค่าในที่อยู่นั้นๆ
	fmt.Println(v)

	changePrice(price)
	fmt.Println("after :", price, addr) // price ไม่เปลี่ยนค่า เนื่องจากตัวแปรถูกเก็บในที่อยู่ที่แตกต่างกัน

	c := &course{"Go", "Nak", 9999} // สามารถเปลี่ยนค่าใน struct ได้ทันที
	c.price = 5599

	fmt.Println("c.price : ", c.price)

	d := discount(c) // หากต้องการเปลี่ยนค่าตัวแปรใน struct ตรงๆต้องใช้ pointer เพื่อระบุที่อยู่ของตัวแปร
	fmt.Println("discount price : ", d)
	fmt.Println("price : ", c.price)

	c1 := new(course) // สร้าง struct zero โดยมี pointer
	fmt.Printf("%#v\n", c1)

	eg := &movie{
		title:       "The Matrix",
		year:        1999,
		rating:      8.7,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Sci-Fi"},
		isSuperHero: true,
	}

	eg.addVote(8)
	fmt.Printf("eg : %#v\n", eg)
	fmt.Printf("votes : %#v\n", eg.votes)
}
