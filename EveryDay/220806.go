package EveryDay

func Func220806() {
	// fmt.Println([...]int{1} == [2]int{1})
	// fmt.Println([]int{1} == []int{1})
}

// go 中不同类型是不能比较的，而数组长度是数组类型的一部分，所以 [...]int{1} 和 [2]int{1} 是两种不同的类型，不能比较；
// 切片是不能比较的；
