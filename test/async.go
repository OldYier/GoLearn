package test

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}
var count = 0

func Async_test() {
	d := [3]string{"绿", "黄", "红"}
	wg.Add(1)
	go loop(d, 0)
	wg.Wait()
	fmt.Println("结束主进程!")
}

func loop(color [3]string, times int) {

	fmt.Println(color[times], time.Now().Second())
	time.Sleep(time.Duration(times+1) * time.Second)
	times++

	times = times % 3
	count++
	fmt.Println("---------", count)
	if count == 10 {
		fmt.Println("第十个了!")
		wg.Done()
	}
	go loop(color, times)
}
