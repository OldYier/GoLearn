package EveryDay

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	// 赋值t=5 -> return赋值 r=t=5 -> defer赋值 t=t+5 = 10 -> 返回r=5
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

//f1 ,f2, f3 分别返回什么?

func Func221130() {
	fmt.Println(f1()) // 1
	fmt.Println(f2()) // 5
	fmt.Println(f3()) // 1
}

// f1() =1，return 把r设成0，然后defer把r改为1 ；
// f2() =5，return 把r设成5，然后defer改的是t，不影响返回值 ；
// f3() =1，return 把r设成1，defer的r是defer的形参,不是return时定义的r,所以defer 的r+5 不影响return的值
