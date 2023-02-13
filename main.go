package main

import (
	"fmt"
	"golearn/EveryDay"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
	"unsafe"
)

func main() {
	// 基础()
	// prConst()
	// prString()
	// 复数()
	// prSlice()
	// prUintptr()
	// 流程控制()
	// prMap()
	// 指定顺序遍历map()
	// a := []int{1, -2, 3, 5, -4, -8, 5}
	// addsum.ThreeSum(a)
	// 多变量赋值()

	EveryDay.Fac230213()

}

type user struct {
	name string
	age  int
}
type Slice2 []int32

func 基础() {
	var a int8 = 10
	var b int8 = 2
	var c = a << b
	var d string = `阿巴阿巴布啦啦
	笑嘻嘻
  	敢敢单单`
	e := 1e9
	fmt.Printf("a的二进制:%b", a)
	fmt.Println("")
	fmt.Printf("b的二进制:%b", b)
	fmt.Println("")
	a = a ^ b
	fmt.Printf("和b异或后a的二进制:%b", a)
	fmt.Println("")
	b = a ^ b
	fmt.Printf("和a异或后b的二进制:%b", b)
	fmt.Println("")
	a = b ^ a
	fmt.Printf("再次和b异或后a的二进制:%b", a)
	fmt.Println("")
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("str := \"c:\\pprof\\main.exe\"")
	fmt.Println(len(d))
	fmt.Println(d)
	fmt.Println("e:", e)
	fmt.Println(reflect.TypeOf(e))
	fmt.Printf("e的内存地址:%p ------ e的类型:%T", &e, e)
	fmt.Println("")
	fmt.Printf("e的二进制:%b", e)
}
func prConst() {
	const (
		a1, a2, a3 = iota, iota + 1, iota + 2
		b1, b2, b3
		c1, c2, c3
		d1, d2, d3
	)
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3)
	fmt.Println(d1, d2, d3)
}
func prString() {
	a := "或,者,这,就是"
	b := []rune(a)
	c := strings.Split(a, ",")
	fmt.Println(c)
	fmt.Printf("C[0]的结构：\t%T\n", c[0])
	fmt.Printf("b[1]的结构:\t%T\n", b[1])
	b[0]++
	fmt.Printf("b[0]++后：\t%v\n", b[0])
	fmt.Println(string(b[0]))
	fmt.Println(b)
	b[0] = '活'
	b[1] = '着'
	fmt.Println(string(b))
}
func 复数() {
	var a complex64 = 2 + 3i

	fmt.Printf("复数a:%G", a)
}
func prSlice() {
	var a [4]string
	a[3] = "刘"
	手机号码 := [3]int64{1: 15616967698, 2: 15573814103}
	书籍 := [4]string{0: "书名：", 1: "福尔摩斯", 2: "作者：", 3: "柯南道尔"}
	fmt.Println(书籍)
	fmt.Println(书籍[0] == 书籍[2])
	fmt.Println(手机号码)
	var b []int = make([]int, 0)
	c := []int{1, 2, 3}
	fmt.Println(b)
	fmt.Println(c)
	d := []string{1: "螺旋丸"}
	fmt.Println(d)
	e := [...]string{0: "时间简史", 1: "百年孤独", 2: "地球往事", 3: "六极物理", 4: "白夜行"}
	fmt.Printf("e的字符输出：%s", e)
	fmt.Println("")
	fmt.Printf("e的指针地址：%p", &e)
	fmt.Println("")
	fmt.Printf("e[0]的指针地址：%p", &e[0])
	fmt.Println("")
	fmt.Printf("e[1]的指针地址：%p\n", &e[1])
	fmt.Printf("截取切片e从2到4 ：%v\n", e[2:4])
	f := []int32{0: 12, 1: 23, 2: 21, 3: 33, 4: 21, 5: 42}
	fmt.Println("--------------------------------")
	fmt.Printf("f的类型：\t%T\n", f)
	fmt.Printf("f对应数组的值：\t%#v\n", f)
	fmt.Printf("f的地址：\t%p\n", &f)
	fmt.Printf("f[0]的地址：\t%v\n", &f[0])
	fmt.Printf("f[1]的地址：\t%p\n", &f[1])
	g := [...]int32{22, 5, 33, 456, 112, 421}
	fmt.Println("--------------------------------")
	fmt.Printf("g的类型:\t%T\n", g)
	fmt.Printf("g的值:\t\t%#v\n", g)
	fmt.Printf("g的地址\t\t%p\n", &g)
	fmt.Printf("g[0]的地址\t%p\n", &g[0])
	fmt.Printf("g[1]的地址\t%p\n", &g[1])
	h := []int{0: 1, 1: 1, 2: 2, 3: 3}
	ss := (*reflect.SliceHeader)(unsafe.Pointer(&h))
	datas := (*[4]int)(unsafe.Pointer(ss.Data))
	fmt.Println("h：")
	fmt.Println(h)
	fmt.Println("ss：")
	fmt.Println(ss)
	fmt.Println("datas：")
	fmt.Println(datas)
	mSlice := make(Slice2, 10, 20)
	mSlice.Append(5)
	fmt.Printf("Append后的mSlice：%v\n", mSlice)
}

// Append func makeslice(et *_type, len, cap int) slice {
//	maxElements := maxSliceCap(et.size)
//	if len < 0 || uintptr(len) > maxElements {
//		panic(errorString("makeslice:len out of range"))
//	}
//	if cap < len || uintptr(cap) > maxElements {
//		panic(errorString("makeslice: cap out of range"))
//	}
//	p := mallocgc(et.size*Uintptr(cap), et, true)
//	return slice{p, len, cap}
//}
func (A1 Slice2) Append(value int32) {
	A2 := append(A1, 2)
	fmt.Println(A1)
	fmt.Println(A2)

	fmt.Printf("A1:\t%v\nA2:\t%v\n", *&A1[9], *&A2[10])
	fmt.Printf("A1:\t%p\nA2:\t%p\n", &A1, &A2)
}
func prUintptr() {
	u := new(user)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	fmt.Println(pName)
	*pName = "张三"

	// 内存偏移
	// 不能分段写，如果分段写可能会导致临时变量被GC，内存变化，操作错误的内存
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 20
	fmt.Println(*u)
}
func 流程控制() {
	a := "上海"
	fmt.Printf("a:\t%#v\n", a)
	a2 := []int32(a)
	fmt.Printf("a2:\t%#v\n", a2)
	a2[1]++
	fmt.Printf("a2[1]:\t%#v\n", a2[1])
	//var x interface{}
	if a2[1] < 28024 {
		fmt.Println("牛逼！")
	} else if a2[1] != 28024 {
		fmt.Printf("a2[1]的值为:\t%#v\n", string(a2[1]))
		a2[1]++
	} else if a2[1] == 28025 {
		fmt.Printf("a2[1]的值为:\t%#v\n", string(a2[1]))
	} else {
		s := a2[1]
		s++
		a2[1]++
		println(s)
		fmt.Printf("a2[1]的值为:\t%#v\n", string(s))
		// fmt.Printf("a2[1]的值为:\t%#v\n", (a2[1]++))
	}
	switch string(a2[1]) {
	case "浹":
		fmt.Println("到了第1层")
		fmt.Println(a2[1])
	case "回":
		fmt.Println("到了第2层")
		fmt.Println(a2[1])
	case "事":
		fmt.Println("到了第3层")
		fmt.Println(a2[1])
	case "扫":
		fmt.Println("到了第4层")
		fmt.Println(a2[1])
	default:
		fmt.Println("默认层")
	}
	//switch x.(type) {
	//case string:
	//	statement(s)
	//}
}
func prMap() {
	map01 := make(map[string]string, 10)
	map01["holmes"] = "就像阳光穿破黑暗"
	map01["oldyier"] = "黎明悄悄来到身边"
	map01["A1"] = "niubi!"
	fmt.Println(map01)
	fmt.Printf("%T\n", map01["holmes"])

	map02 := map[string]int{
		"appSum":   20,
		"peSum ":   20,
		"unappSum": 30,
		"niuSum":   15,
	}
	fmt.Printf("%v\n", map02)

	value, ok := map02["app1Sum"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在这个键")
	}

	for key, v := range map02 {
		fmt.Println(key, "\t的值为:\t", v)
	}

	fmt.Println("遍历键值对的键")

	for key := range map02 {
		fmt.Println("map02的key: ", key)
	}

	// 删除map中的键值对
	delete(map02, "appSum")

	fmt.Println("删除后的map02:")
	for key := range map02 {
		fmt.Println("map02的key: ", key)
	}
}

func 指定顺序遍历map() {
	fmt.Println("time.Now()的值: ", time.Now())
	fmt.Println("time.Now().UnixNano()的值: ", time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	var randMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		randMap[key] = value

		var keys = make([]string, 0, 200)
		for key := range randMap {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			fmt.Println(key, randMap[key])
		}

	}
}

func 多变量赋值() {
	var a int8 = 1
	var b int8 = 2
	var c int8 = 3
	a, b, c = 4, c, b
	print("a, b, c:", a, b, c)
}
