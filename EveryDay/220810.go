package EveryDay

import "fmt"

func Func220810(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var Func220810 func()

	defer Func220810()
	Func220810 = func() {
		r += 2
	}
	return n + 1
}

func External() string {
	fmt.Println(Func220810(3))
	return "success"
}
