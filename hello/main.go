package main

/*
123
++ package dùng để nhóm 1 hoặc nhiều tập tin có liên quan đến nhau
++ tên package sử dụng in thường, định nghĩa ở đầu chương trình
++ tập để chạy chương trình thì tên là package main, đồn thời phải khởi tạp func main
	trình Go compiler sẽ tìm func main để run
*/

/*
++ import package để sử dụng 1 package khác
++ "fmt" : package formated input/output xử lý nhập xuất
*/

import "fmt"

func main() {
	fmt.Println("hello duong")

	fmt.Println(sum(5))
}

// khởi tạo 1 hàm Sum, truyền vào number (biến int), trả về int
func sum(number int) int {
	s := 0                        // khai báo biến
	for i := 0; i < number; i++ { // vòng lặp
		s += i
	}

	return s
}
