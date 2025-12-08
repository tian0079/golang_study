package main

import (
	"fmt"
	"golang_study/task1"
	"golang_study/task2"
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
	sum, ok := task1.TwoSum(nums, target)
	if ok {
		fmt.Println("两数之和", sum)
	}

	//指针
	number := 5
	fmt.Println("指针", task2.ChangeNumber(&number))
}
