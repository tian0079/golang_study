package task2

import (
	"fmt"
	"sync"
)

func GoroutineFunc() {
	var wg sync.WaitGroup
	var slices []int

	for i := 1; i <= 10; i++ {
		slices = append(slices, i)
	}
	wg.Add(1)
	go printOddNumber(slices, &wg)
	wg.Add(1)
	go printEvenNumber(slices, &wg)
	wg.Wait()
}

func printOddNumber(numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 != 0 {
			fmt.Println("奇数", numbers[i])
		}
	}
}
func printEvenNumber(numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			fmt.Println("偶数", numbers[i])
		}
	}
}
