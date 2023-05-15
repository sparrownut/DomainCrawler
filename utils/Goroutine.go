package utils

import (
	"sync"
	"time"
)

type Task func()

func RunWithConcurrency(tasks []Task, concurrency int) {
	// 创建一个 wait group
	var wg sync.WaitGroup

	// 创建一个 buffer chan，这个 chan 的大小就是并发的数量
	sem := make(chan bool, concurrency)

	// 遍历所有的任务
	for _, task := range tasks {
		sem <- true // 向 sem 发送一个信号，如果 sem 满了，那么这行代码会阻塞，直到 sem 中有空位
		wg.Add(1)   // 增加一个任务

		// 启动一个协程
		go func(task Task) {
			time.Sleep(time.Millisecond * 100)
			defer wg.Done()          // 任务完成，通知 wait group
			defer func() { <-sem }() // 任务完成，从 sem 中释放一个位置

			task() // 执行任务
		}(task)
	}

	wg.Wait() // 等待所有任务完成
}
