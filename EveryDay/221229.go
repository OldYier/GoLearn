package EveryDay

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

func Func221229() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	fmt.Println("原始切片:", slice)
	change(slice...)
	fmt.Println("修改切片:", slice)
	fmt.Println("临时切片:", slice[0:2])
	change(slice[0:2]...)
	fmt.Println("最终切片:", slice)
	// Go 提供的语法糖...，可以将 slice 传进可变函数，不会创建新的切片。
	// 第一次调用 change() 时，append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变；
	// 第二次调用 change() 函数时，使用了操作符[i,j]获得一个新的切片，假定为 slice1，它的底层数组和原切片底层数组是重合的，
	// 不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。
}
