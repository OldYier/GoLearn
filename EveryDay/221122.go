package EveryDay

import "fmt"

func Func221122() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Printf("a : %v \t", a)
	fmt.Printf("b : %v \t", b)
	fmt.Printf("c : %v \t\n", c)
	fmt.Printf("a.len : %v \n", len(a))
	fmt.Printf("b.len : %v \n", len(b))
	fmt.Printf("c.len : %v \n", len(c))
	fmt.Printf("a.cap : %v \n", cap(a))
	fmt.Printf("b.cap : %v \n", cap(b))
	fmt.Printf("c.cap : %v \n", cap(c))
}

// 切片 a、b、c 的长度和容量分别是多少？
// 参考答案及解析：a、b、c 的长度和容量分别是 0 3、2 3、1 2。
// 知识点：数组或切片的截取操作。截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，
// 假设截取对象的底层数组长度为 l。在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。
// 操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。
