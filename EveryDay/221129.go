package EveryDay

import "fmt"

type A1129 interface {
	ShowA() int
}

type B1129 interface {
	ShowB() int
}

type Work1129 struct {
	i int
}

func (w Work1129) ShowA() int {
	return w.i + 10
}

func (w Work1129) ShowB() int {
	return w.i + 20
}

func Func221129() {
	var a A1129 = Work1129{3}
	s := a.(Work1129)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
