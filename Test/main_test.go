package main

import "testing"

func TestAdd(t *testing.T) { // ต้องขึ้นต้นด้วย Test
	t.Run("when input 1 and 2 should return 3", func(t *testing.T) { // ระบุข้อความ test เพิ่มเติม
		if Add(1, 2) != 3 {
			t.Error("1 + 2 should be 3")
		}
	})

	t.Run("when input 2 and 2 should return 4", func(t *testing.T) { // ระบุข้อความ test เพิ่มเติม
		if Add(2, 2) != 4 {
			t.Error("2 + 2 should be 4")
		}
	})
}

// go test . // run test
// go test -v . // -v คือ verbose ไว้ระบุรายละเอียดเพิ่มเติม
// go test -v ./... // ไว้ run ทุก package
// ต้อง go mod init ก่อน
