package main

import "fmt"

func main() {
	//khai bao array
	var myArray [4]int
	fmt.Println(myArray) // [0 0 0 0]
	// phần tử mảng chưa đc gán giá trị sẽ đc gán mặc định 0

	myArray[0] = 11
	myArray[1] = 23
	fmt.Println(myArray) // [11 23 0 0]

	//khai báo có khởi tạo giá trị
	arrays := [3]int{1, 2, 3}
	//  var arrays = [3]int {1,2,3}
	fmt.Println(arrays)

	// ko thể sử dụng phần tử ngoài khai báo, ví dụ arrays[3]

	//check size mảng
	fmt.Println(len(myArray))

	// khởi tạo ko khai báo size
	arrays1 := [...]string{"nguyen", "thai", "duong"}
	fmt.Println(arrays1)

	// có thể gán 2 mảng cho nhau
	copyArrays1 := arrays1
	fmt.Println(copyArrays1)

	// mảng copyArrays1 và arrays1 là hoàn toàn độc lập với nhau
	// ko giống trong C++ là 2 mảng cùng trỏ về 1 vùng nhớ

	// duyệt mảng, dùng for

	// cách 1
	for i := 0; i < len(arrays); i++ {
		//
	}

	// cách 2
	for index, value := range arrays1 {
		fmt.Printf("i = %d value = %s", index, value)
		fmt.Println()
	}

	// nếu chỉ làm việc với 1 trong 2 phần tử index, value thay phần tử đó bằng '_'
	for _, value := range arrays1 {
		fmt.Printf("value = %s", value)
		fmt.Println()
	}

	//mảng 2 chiều
	matrix := [4][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println(matrix)

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j], "    ")
		}
		fmt.Println()
	}

}
