package main

import (
	"fmt"
)

func main() {
	number := 10

	if number == 10 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	fmt.Println("==================")

	// khởi tạo kết hợp so sánh
	if a := 100; a > 100 {
		fmt.Println("a>100")
	} else {
		fmt.Println("a <= 100")
	}

	fmt.Println("==================")

	// lưu ý, golang ko nhận else xuống dòng, phải viết else ngay sau } của if

	number1 := 10

	switch number1 {
	case 1:
		fmt.Println("1")
	case 4, 5, 9: // bắt nhiều giá trị
		fmt.Println("4")
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("high")
	default:
		fmt.Println("unknow")
	}

	fmt.Println("==================")

	switch {
	case number1 > 100:
		fmt.Println("high")
	case number1 <= 10:
		fmt.Println("low")
	}

	fmt.Println("==================")

	// switch : fallthrough
	switch number1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 4:
		fmt.Println("4")
		fallthrough
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("high")
		fallthrough
	default:
		fmt.Println("unknow")
	}

	// fallthrough chạy hết các giá trị kể từ case đúng

}
