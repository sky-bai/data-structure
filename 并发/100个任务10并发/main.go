package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 100
	max := 10
	var wg sync.WaitGroup
	ch := make(chan struct{}, max)
	defer close(ch)
	for i := 0; i < count; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go task(ch, wg)
	}
	wg.Wait()

}

func task(ch chan struct{}, wg sync.WaitGroup) {

	fmt.Println("任务")
	defer wg.Done()
	<-ch

}

// 用waitgroup去创建子任务 再创建最大容量的chan 每个任务通过waitgroup创建 在创建的时候往chan里面发送数据 如果此时chan里面是满的就阻塞，说明
// 此时已达到最大的并发数 ，然后每一个任务在完成时 从chan里面取数据 表示该任务已完成
