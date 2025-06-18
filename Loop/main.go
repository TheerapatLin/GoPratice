package main

import "fmt"

func main() {

	sum := 1

	for i := 0; i < 5; i++ {
		// sum := 0 //sum ถูกประกาศใหม่เสมอ
		sum += sum
		fmt.Println("i : ", i, "sum : ", sum)
	}
	fmt.Println("sum loop1 done : ", sum)

	sum1 := 1

	for sum1 < 5 { //while loop
		sum1 += sum1
		fmt.Println("sum1 : ", sum1)
	}
	fmt.Println("sum1 loop2 done : ", sum1)

	genres := [...]string{"action", "horror", "comedy"}
	for i := 0; i < len(genres); i++ {
		fmt.Println(genres[i])
	}

	for i := range genres {
		fmt.Println(genres[i])
	}

	for _, val := range genres { // ตัวแปรแรก จะแทน index เสมอหากมีการประกาศ 2 ตัวแปร, _ ประกาศตัวแปรแต่ไม่ได้ใช้
		fmt.Println("value : ", val)
	}

}
