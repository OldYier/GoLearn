package EveryDay

import "fmt"

// 不能通过,new 产生的是指针类型,append要的是值类型
func Func220909() {
	list := new([]int)
	// list = append(list, 1)
	fmt.Println(list)
}
