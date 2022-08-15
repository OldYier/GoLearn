package EveryDay

import "fmt"

func Func220808() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
		fmt.Println("V:", v, " I:", i)
	}
}
