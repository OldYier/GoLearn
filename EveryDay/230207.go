package EveryDay

import (
	"fmt"
)

type People07 interface {
	Show()
}

type Student07 struct{}

func (stu *Student07) Show() {

}

func live() People07 {
	var stu *Student07
	return stu
}

func Func230207() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
