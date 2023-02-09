package EveryDay

import "fmt"

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func Func221128() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

// func increaseA() int {，返回值i=0的时候已经绑定到返回值里里，defer改i没用了。func increaseB() (r int) {先把返回变量r设为0，defer又把r改为1.
// return语句是把return后面的值赋给返回值，但由于increaseB中有return变量r，所以在defer里还可以对变量r进行再更新，所以返回的r为1
