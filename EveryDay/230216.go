package EveryDay

import "fmt"

type person struct {
	name string
}

func Func230216() {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}
