package EveryDay

import "fmt"

func Func220825() {
	Func220825_1()
	Func220825_2()
}

func Func220825_1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// 2.
func Func220825_2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}
