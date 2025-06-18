package main

import "fmt"

func main() {
	defer fmt.Println("a") // defer จะทำการรันคำสั่งบรรทัดนี้ทีหลัง
	defer fmt.Println("b")
	fmt.Println("c")

	fmt.Println("Counting...")
	for i := 0; i < 10; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
	fmt.Println("Done")

}
