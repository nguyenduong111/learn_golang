package main

import "fmt"

/*
+ slice tham chiếu đến 1 mảng, mô tả 1 phần hoặc toàn bộ mảng
+ slice có kích thước động nên ko phải khai báo size khi khởi tạo

*/

func main() {
	// khai báo silce
	var slice []int
	fmt.Println(slice)

	//khai bao va khoi tao
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	// tham chiếu đến mảng
	var array = [4]int{1, 2, 3, 4}
	slice2 := array[1:3] // gán vào slice2 từ array[1] -> array[3-1]
	fmt.Println(slice2)

	/*
		slice2 := array[:]     // lấy toàn bộ array gán vào slice2
		slice2 := array[1:]    // lấy từ array[1] đến hết
	*/

	//slice tham chiếu đến mảng, nên khi đó, các thay đổi trong slice sẽ thay đổi array đc tham chiếu
	/*
		var array = [4]int {1,2,3,4}
		slice2 := array[:]
		slice2[0] = 55

		fmt.Println(slice2)     // [55, 2, 3, 4]
		fmt.Println(array)      // [55, 2, 3, 4]

	*/

	/*
		phân biệt len và cap trong slice
		+ len là số phần tử trong slice
			fmt.Println(len(slice2))      // 4

		+ cap (capacity) : số lượng phần tử tính từ vị trí bắt đầu gán array vào slice đến cuối
			nếu gọi start = index của phần tử đầu tiên gán vào slice
			cap(slice) = len(array) - start

	*/

	// một số hàm lm việc với slice

	// make : khởi tạo slice với len và cap cung cấp
	//slice3 := make([]int, 2, 5)
	// len = 2, cap = 5

	//slice4 := make([]int, 2)
	// len = 2, cap = 2

	// append : thêm phần tử vào slice
	var slice5 []int
	slice5 = append(slice5, 11)

	// copy (return số phần tử đc copy):
	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 2)

	number := copy(dest, src)
	fmt.Println(number) // 2
	fmt.Println(dest)   // [A, B]

	// nối 2 slice
	//des = append(slice1[:], slice2[:])

}
