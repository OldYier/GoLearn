package EveryDay

import "fmt"

//func Func221123() {
//	var m map[string]int //A
//	m["a"] = 1
//	if v := m["b"]; v != nil { //B
//		fmt.Println(v)
//	}
//}

func Func221123G() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}
}

// 在A处只声明了map m ,并没有分配内存空间,不能直接赋值,需要会用make(),都提倡使用make() 或者字面量的方式直接初始化map.
// B 处,v,k := m["b"] 当key为b的元素不存在的时候,v会返回值类型对应的零值,k返回false.
