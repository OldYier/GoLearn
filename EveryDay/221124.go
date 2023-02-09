package EveryDay

type A interface {
	showA() int
}

type B interface {
	showB() int
}

type Work2 struct {
	i int
}

func (w Work2) showA() int {
	return w.i + 10
}

func (w Work2) showB() int {
	return w.i + 20
}

func Func221124() {
	// c := Work2{3}
	// var a A = c
	// var b B = c
	// fmt.Println(a.showB()) // 13
	// fmt.Println(b.showA()) // 5
}

// 接口的静态类型。a、b 具有相同的动态类型和动态值，分别是结构体 work 和 {3}；a 的静态类型是 A，b 的静态类型是 B，接口 A 不包括方法 ShowB()，接口 B 也不包括方法 ShowA()，编译报错。
