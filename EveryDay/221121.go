package EveryDay

import (
	"fmt"
)

type interA interface {
	ShowA() int
}

type interB interface {
	ShowB() int
	ShowA() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func Func221121() {
	c := Work{3}
	var a interA = c
	var b interB = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}

// 接口。一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。
