package EveryDay

import (
	"fmt"
	"time"
)

func say(s string, t int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Fac230213() {
	go say("world", 101)
	say("hello", 100)

	// ch := make(chan int, 1000)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		a, ok := <-ch
	// 		if !ok {
	// 			fmt.Println("close")
	// 			return
	// 		}
	// 		fmt.Println("a: ", a)
	// 	}
	// }()

	// fmt.Println("ok")
	// time.Sleep(time.Second * 1)
	// close(ch)
}
