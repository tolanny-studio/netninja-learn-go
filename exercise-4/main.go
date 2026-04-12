package main

import (
	"fmt"
)

func main() {
	// array and slices

	// array
	// var nums [3]int = [3]int{2, 6, 8}

	var nums = [6]int{2, 6, 8}
	fmt.Println(nums)

	// two-dimensional array

	two_dim_array := [...][5]int{{1, 4, 7}, {9, 2, 3, 6}}

	for row := range two_dim_array {
		for col := range two_dim_array[row] {
			fmt.Println(two_dim_array[row][col])
		}
	}

	// slices
	slice_1 := []int{3, 5, 7, 9}
	slice_1 = append(slice_1, 4, 5, 8)

	slice_2 := slice_1[2:5]
	fmt.Println(slice_1)
	fmt.Println(slice_2)
	fmt.Println(len(slice_2))

}
