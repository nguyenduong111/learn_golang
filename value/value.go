package main

import "fmt"

func main() {

	// 1
	var number int
	number = 10
	fmt.Println(number)

	// 2
	var number1 int = 11
	fmt.Println(number1)

	// 3: type Inference tự động ép kiểu
	var email = "duong@gmail.com" // string
	fmt.Println(email)

	// 4 khai bao nhieu bien cung kieu
	var a, b int
	a = 1
	b = 2
	fmt.Println(a)
	fmt.Println(b)

	var a2, b2 = 12, 13
	fmt.Println(a2)
	fmt.Println(b2)

	// 5 khai bao nhieu bien khac kieu
	var (
		name    string
		address string
		age     int
	)

	name = "nguyen duong"
	address = "nam dinh"
	age = 21

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)

	var (
		name1    string = "nhung"
		address1 string = "hh"
		age1     int    = 19
	)

	fmt.Println(name1)
	fmt.Println(address1)
	fmt.Println(age1)

	var name2, address2, age2 = "name2", "add2", 1

	fmt.Println(name2)
	fmt.Println(address2)
	fmt.Println(age2)

	// 6 short hand
	year := "hihihi"

	fmt.Println(year)
}
