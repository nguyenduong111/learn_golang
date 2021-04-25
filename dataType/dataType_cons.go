package main

//kieu dun lieu

import "fmt"

func main() {

	// bool
	var myBool bool = true // false
	fmt.Println(myBool)

	fmt.Println("==============")

	// string
	var myString string = "abcdef"
	fmt.Println(myString)

	fmt.Println("==============")

	// int
	var myInt int = 11
	fmt.Println(myInt)

	// int8, 16, 32, 64 : so bit max bieu dien int

	// int8 : -2^7 .... 2^7
	// int16: -2^15.... 2^15
	// int32: -2^31.... 2^31

	fmt.Println("==============")

	// uint : nguyen duong (8,16,32,64)
	var mtUint uint = 10
	fmt.Println(mtUint)

	fmt.Println("==============")

	// byte: có bản chất là uint8
	var myByte byte = 111
	fmt.Println(myByte)

	fmt.Printf("%T", myByte) // %T : Tyte data

	fmt.Println()
	fmt.Println("==============")

	var a byte = 'A' // byte voi char se return ascii
	fmt.Println(a)   // 65

	var b string = "A"
	fmt.Printf(b) // A

	fmt.Println("==============")

	// float32 or float64
	var myFloat float32 = 10.11
	fmt.Println(myFloat)

	fmt.Println("==============")

	// complex64 or complex128: so phuc
	var z1 complex64 = 10 + 3i
	fmt.Println(z1)

	fmt.Println("==============")

	// rune
	var str string = "dương"

	for i := 0; i < len(str); i++ { // len(stringName) độ dài chuỗi
		fmt.Printf("%c", str[i])
	}

	// in string tiếng anh thì ok, tiếng việt sẽ lỗi
	// ép kiểu string sang rune

	fmt.Println()
	myRune := []rune(str) // ép từ string sang rune
	for i := 0; i < len(myRune); i++ {
		fmt.Printf("%c", myRune[i])
	}
	fmt.Println()

	// string là 1 chuỗi các bytes
	// rune (đại diện cho int32) 1 rune có thẻ chứa nhiều bytes. có thẻ coi 1 chuỗi là dãy rune
	// có thể chuyển đổi từ string sang rune và ngược lại
	// dùng rune để in ra các kí tự đặc biệt

	// xem thêm các định dạng xuất của thư viện fmt
	// vào fmt golang format

	fmt.Println("==============")

	// ep kieu

	// ko tự động ép kiểu giá trị
	// tính toán cùng kiểu dữ liệu

	var a1 int = 1
	var b1 float32 = float32(a1)
	fmt.Println(b1)
	fmt.Println(a1 + int(b1))

	fmt.Println("==============")

	// const

	const PI = 3.14
	fmt.Println(PI)

}
