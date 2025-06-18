package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

type movie struct {
	title string
	year int
	rating float32
	genres []string
	isSuperHero bool
}

func (m movie) info() {
	fmt.Printf("%s (%d) - %.2f\n",m.title,m.year,m.rating)
	fmt.Println("Genres : ")
	for _,g := range m.genres {
		fmt.Printf("\t%s\n", g)		// \t คือการ tab
	}
}

// function
// func discount(c course) int { // รับค่าชนิด course เก็บในตัวแปร c
// 	p := c.price - 599
// 	fmt.Println("Discounted price is", p)
// 	return p
// }

// method : property ที่เป็น func
func (c course) discount(d int) int { 
	p := c.price - d
	fmt.Println("Discounted price is", p)
	return p
}

func (c course) info() {	// (c course) คือ receiver
	fmt.Println("Name : ", c.name)
	fmt.Println("Instructor : ", c.instructor)
	fmt.Println("Price : ", c.price)
}

func main() {
	c := course{
		name:       "Go",
		instructor: "Mihalis",
		price:      2000,
	}
	fmt.Printf("%#v\n", c)

	// d := discount(c)	// เรียก func
	d := c.discount(499) // เรียก method
	fmt.Println("discount price : ", d)
	c.info()

	ae := movie{
		title: "Avengers: Endgame",
		year: 2019,
		rating: 8.4,
		genres: []string{"Action", "Adventure", "Sci-Fi"},
		isSuperHero: true,
	}

	ae.info()
 
	}
	

