package main

// คำสั่ง go mod tidy จะตามหา package ที่จะถูก import ใน internet
import (
	"fmt" // import build-in func

	"github.com/TheerapatLin/cinema/movie"  // import package ที่สร้างขึ้นเอง
	"github.com/TheerapatLin/cinema/ticket" // import package ที่สร้างขึ้นเอง
)

func init() { // func ที่เรียกใช้งานโดยอัตโนมัติเมื่อถูกเรียกใช้งาน package นั้น // เรียก init ตามลำดับของการ import
	fmt.Println("init : main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println("movieName : ", movieName)

	movie.Review(movieName, 8.5)

	ticket.BuyTicket(movieName)
}
