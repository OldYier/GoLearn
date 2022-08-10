package EveryDay

import "fmt"

func Func220808() string {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
		fmt.Printf("v:%v \t\ti :%v\n", v, i)
	}
	return "success"
}
