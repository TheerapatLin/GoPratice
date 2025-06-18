package ticket

import "fmt"

func init() {
	fmt.Println("init : ticket")
}

func BuyTicket(movie string) { // camel case ขึ้นต้นด้วยตัวเล็ก
	fmt.Printf("I bought tickets for %s\n", movie) // ใช้ชื่อ package . ด้วยชื่อ func นั้นๆ (fmt.Println())
}
