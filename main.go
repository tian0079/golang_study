package main

import (
	"fmt"
	"golang_study/task1"
	"golang_study/task2"
)

func main() {

	fmt.Println("-------------------------task1 两数之和 START------------------------------")
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

	fmt.Println("-------------------------task2 两数之和 END--------------------------------")

	fmt.Println()

	fmt.Println("-------------------------task2 指针START--------------------------------")
	//指针
	number, number1 := 5, 5
	task2.ChangeNumber(&number)
	fmt.Println("指针 题目1:", number)

	task2.Multiplication(&number1)
	fmt.Println("指针 题目2:", number1)
	fmt.Println("-------------------------task2 指针END--------------------------------")

	fmt.Println()

	fmt.Println("-------------------------task2 Goroutine START------------------------")
	fmt.Println("题目1：")
	task2.GoroutineFunc()

	fmt.Println("题目2：")

	scheduler := &task2.TaskScheduling{}

	scheduler.AddTask(func() (string, interface{}) {
		return "任务1", "结果1"
	})
	scheduler.AddTask(func() (string, interface{}) {
		return "任务2", "结果2"
	})
	scheduler.AddTask(func() (string, interface{}) {
		return "任务3", "结果3"
	})

	scheduler.Run()

	fmt.Println("-------------------------task2 Goroutine END--------------------------")
	fmt.Println()
	fmt.Println("-------------------------task2 面相对象 START------------------------")
	fmt.Println("题目1")
	Rectangle := &task2.Rectangle{}
	Rectangle.Perimeter()
	Rectangle.Area()

	fmt.Println("题目2")
	Employee := &task2.Employee{}
	Employee.EmployeeID = 123
	Employee.Age = 23
	Employee.Name = "测试姓名"

	Employee.PrintInfo()

	fmt.Println("-------------------------task2 面相对象 END--------------------------")
	fmt.Println()

	fmt.Println("-------------------------task2 Channel START--------------------------")
	fmt.Println("题目1")
	task2.RunPrintNumberFunc()
	fmt.Println()
	fmt.Println("题目2")
	//带缓冲的
	task2.PrintChannelBuffer()

	fmt.Println("-------------------------task2 Channel END----------------------------")
	fmt.Println()
	fmt.Println("-------------------------task2 锁机制 START--------------------------")
	fmt.Println("题目1")
	lockShared := task2.LockSharedCounter{}
	lockShared.SharedCounter()
	fmt.Println()
	fmt.Println("题目2")
	lockShared = task2.LockSharedCounter{}
	lockShared.ActionCount()

	fmt.Println("-------------------------task2 锁机制 END----------------------------")
}
