package EveryDay

import "fmt"

type Student02 struct {
	Name string
}

var list map[string]*Student02

func Func230202() {

	list = make(map[string]*Student02)

	student := Student02{"Aceld"}

	list["student"] = &student
	list["student"].Name = "LDB"

	fmt.Println(list["student"])
}
