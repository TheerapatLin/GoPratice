package movie

import "fmt"

func init() {
	fmt.Println("init : movie")
}

// การตั้งชื่อ func หากตัวอักษรตัวแรกขึ้นต้นด้วยตัวพิมพ์ใหญ่ func นั้นจะถือว่าเป็น Public func กลับกันหากเป็นตัวพิมพ์เล็ก จะถือว่าเป็น Private func
// private func สามารถมองเห็นได้แค่ใน package นั้นๆ
func Review(name string, rating float64) { // Pascal case ขึ้นต้นด้วยตัวใหญ่
	fmt.Printf("I reviewed %s and it's rating is %f\n", name, rating)
}
