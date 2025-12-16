package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type LockSharedCounter struct {
	count int64
	mu    sync.Mutex
}

func (lockSharedCounter *LockSharedCounter) SharedCounter() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < 1000; i++ {
				lockSharedCounter.increment()
			}

		}()
	}
	wg.Wait()
	fmt.Println("10个协程 1000次递增结果：", lockSharedCounter.count)
}

func (lockSharedCounter *LockSharedCounter) increment() {
	lockSharedCounter.mu.Lock()
	lockSharedCounter.count++
	lockSharedCounter.mu.Unlock()
}

func (lockSharedCounter *LockSharedCounter) ActionCount() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				lockSharedCounter.incrementAtomic()
			}
		}()
	}
	wg.Wait()
	fmt.Println("10个协程 1000次递增结果：", lockSharedCounter.count)
}

func (lockSharedCounter *LockSharedCounter) incrementAtomic() {
	atomic.AddInt64((*int64)(&lockSharedCounter.count), int64(1))
}
