package main

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

func add(a ...int) int {
	var sum int
	for _, i := range a {
		sum += i
	}
	return sum
}

func TestDemo(t *testing.T) {
	// func add(a ...int) int {
	// fmt.Println(add([]int{1, 2}))
	// fmt.Println(add([]int{1, 3, 7}...))
	// func Println(a ...any) (n int, err error) {
	// fmt.Println([]string{"1", "2"})
	// fmt.Println([]any{"1", "2"}...)

	var ch = make(chan int, 1)
	ch <- 100
	close(ch)
	<-ch
	fmt.Println(<-ch)
}

func TestShuChuYaZhi(t *testing.T) {
	n := 10
	my := []int{10000, 1000, 1000, 1111, 11111, 11212, 11213, 1111, 11111, 11212}
	enemy := []int{10000, 1000, 1000, 1111, 11111, 112120, 11213, 1111, 11111, 11212}
	dis := make([]int, n)

	for i := range my {
		dis[i] = enemy[i] - my[i]
	}

	var maxKr int
	for i, v := range dis {
		if v > 0 {
			if v > maxKr*(i+1) {
				maxKr = (v+1)/(i+1) + 1

				fmt.Println(maxKr*(i+1), v)
			}
		}
	}
	fmt.Println(maxKr)
}
