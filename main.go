package main

import (
	"fmt"
	"golang_study/task1"
)

func main() {

	// nums := []int{
	// 	2, 7, 11, 15,
	// }
	nums := []int{
		3, 2, 4,
	}
	// nums = []int{
	// 	3, 3,
	// }
	target := 6
	fmt.Println(task1.TwoSum(nums, target))
}
