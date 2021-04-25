package main

import (
	"fmt"
)

func ham1() { // void , no params
	fmt.Println("hello")
}

func ham2(name string) { // void, 1 params
	fmt.Println(name)
}

func ham3(name string) string { // return string, 1 params
	kq := name + " " + "nguyen"
	return kq
}

// trong go, 1 func có thể return về nhiều giá trị
func ham4(w, h int) (int, int) { // 2 return. 2 params
	return w, h
}

// name return value
func ham5(w, h int) (wi int, hei int, isS bool) {
	isS = w == h
	wi = w
	hei = h
	return
}

func main() {
	name := "duong"

	ham1()

	ham2(name)

	fmt.Println(ham3(name))

	w, h := ham4(100, 300)
	fmt.Println(w)
	fmt.Println(h)

	w1, h1, isS := ham5(1, 2)
	if isS {
		fmt.Println("fall")
	} else {
		fmt.Println(w1)
		fmt.Println(h1)
	}
}
