package task2

import (
	"fmt"
	"sync"
	"time"
)

type Task func() (string, interface{})
type TaskScheduling struct {
	tasks []Task
	wg    sync.WaitGroup
}

func (ts *TaskScheduling) AddTask(task Task) {
	ts.tasks = append(ts.tasks, task)
}

func (ts *TaskScheduling) Run() {
	for _, task := range ts.tasks {
		ts.wg.Add(1)
		go func(t Task) {
			defer ts.wg.Done()
			start := time.Now()
			name, result := t()
			fmt.Printf("Task %v result %v 运行时间：%v\n", name, result, time.Since(start))
		}(task)
	}

	ts.wg.Wait()
}
