package go_edu

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Task string

const (
	K         = 10
	workToken = "inWork"
)

// 调用方法 goroutine 睡眠 2 秒，模拟任务执行
func execute(task Task) {
	// fmt.Println(string(task))
	time.Sleep(time.Second * 2)
}

// 可以通过 buf-channel 非常优雅的实现
func TestDay06(t *testing.T) {
	tasks := []Task{
		"this task 01",
		"this task 02",
		"this task 03",
		"this task 04",
		"this task 05",
		"this task 06",
		"this task 07",
		"this task 08",
		"this task 09",
		"this task 10",
		"this task 11",
		"this task 12",
		"this task 13",
		"this task 14",
		"this task 15",
	}

	var (
		wg = sync.WaitGroup{}
		ch = make(chan string, K)
	)

	defer close(ch)

	for idx, task := range tasks {
		wg.Add(1)
		go func(id int, task Task) {
			// 任务开始前放入 token，如果 buffer 区已满则会阻塞
			ch <- workToken
			t := time.Now()
			execute(task)
			fmt.Printf("「%v」号任务执行, 用时：%v ms , 当前并发 execute goroutine 数: %v\n",
				id,
				time.Now().UnixMilli()-t.UnixMilli(), len(ch))
			<-ch
			wg.Done()
		}(idx, task)
	}
	wg.Wait()
}

func TestDay04(t *testing.T) {
	// 1、
	// 通过查阅文档，发现 go test 指令执行中，为了加快编译速度，不同包的编译是并发执行的；
	// 但 test 内方法执行保障其串行化，才能尽可能的保证执行过程中不受其他任务的影响。
	// 调用 t.Parallel() 可提供一个信号量，执行时 该 test 下的任务可和其他被 Parallel 标记的任务并发执行。

	// 2、大文件分段读入

}

