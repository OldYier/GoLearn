package test

import (
	"fmt"
	"time"
)

var myCount = 0

func PrintFibonacci() {
	// for i := 1; i < 94; i++ {
	fmt.Println("开始时间:", time.Now().UnixNano())
	// res := fibonacci(50)
	now := time.Now()
	res := fibonacci2(150)
	fmt.Printf("fibonacci %d = %d \n", 150, res)
	fmt.Println("结束时间1:", time.Now().UnixNano())
	fmt.Println("time elapse:", time.Since(now))
	res2 := fibonacci2(150)
	fmt.Printf("fibonacci %d = %d \n", 150, res2)
	// fmt.Printf("myCount\t = %d \n", myCount)
	fmt.Println("结束时间2:", time.Now().UnixNano())
	fmt.Println("time elapse:", time.Since(now))
	// }
}

func fibonacci(i int) (res int) {

	if i <= 2 {
		res = 1
	} else {
		res = fibonacci(i-1) + fibonacci(i-2)
	}
	return
}

func fibonacci2(N int) int {
	cur, next := 0, 1
	for i := 0; i < N; i++ {
		cur, next = next, cur+next
	}
	return cur
}

func fibonacci3(N int) int {
	if N < 2 {
		return N
	}
	fibMap := make(map[int]int)
	fibMap[0] = 0
	fibMap[1] = 1
	for i := 2; i <= N; i++ {
		fibMap[i] = fibMap[i-1] + fibMap[i-2]
	}
	return fibMap[N]
}
