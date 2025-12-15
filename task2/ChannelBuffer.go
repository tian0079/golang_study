package task2

import (
	"fmt"
	"sync"
)

//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

func PrintChannelBuffer() {
	wg := sync.WaitGroup{}
	intNum := 100
	intChan := make(chan int, intNum)
	wg.Add(1)
	go func() {
		defer wg.Done()
		beginGenerateNumbers(intChan, intNum)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		ConsumeNumbers(intChan)
	}()

	wg.Wait()
}

func beginGenerateNumbers(intChan chan<- int, ints int) {
	for i := 1; i <= ints; i++ {
		intChan <- i
	}
	close(intChan)
}

func ConsumeNumbers(intChan <-chan int) {
	for num := range intChan {
		fmt.Println(num)
	}

}
