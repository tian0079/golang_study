package task2

import (
	"fmt"
	"sync"
)

func RunPrintNumberFunc() {
	var ints int = 10
	wg := sync.WaitGroup{}
	intChan := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		CreateNumber(ints, intChan)

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		PrintNumber(intChan)
	}()
	wg.Wait()
}
func CreateNumber(ints int, intChan chan<- int) {
	for i := 1; i <= ints; i++ {
		intChan <- i
	}
	close(intChan)
}

func PrintNumber(intChan <-chan int) {
	for num := range intChan {
		fmt.Println(num)
	}
}
