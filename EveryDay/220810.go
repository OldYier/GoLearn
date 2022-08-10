package EveryDay

import "fmt"

func Func220810(n int) (r int) {
	// 执行延迟函数
	// Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。
	// 通常来说，不应该对进入 panic 宕机的程序做任何处理，但有时，需要我们可以从宕机中恢复，至少我们可以在程序崩溃前，做一些操作，举个例子，当 web 服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接关闭，如果不做任何处理，会使得客户端一直处于等待状态，如果 web 服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试。提示
	// 在其他语言里，宕机往往以异常的形式存在，底层抛出异常，上层逻辑通过 try/catch 机制捕获异常，没有被捕获的严重异常会导致宕机，捕获的异常可以被忽略，让代码继续运行。
	// Go语言没有异常系统，其使用 panic 触发宕机类似于其他语言的抛出异常，recover 的宕机恢复机制就对应其他语言中的 try/catch 机制。
	defer func() {
		r += n
		recover()
	}()

	// 定义了一个方法
	var Func220810 func()
	// Func220810 未初始化 不会执行
	defer Func220810()

	// 方法初始化 ,但是没有执行
	Func220810 = func() {
		r += 2
	}
	// 所以先return  4
	return n + 1
	// 执行顺序 :
	// return r = 4
	// 延迟方法 func r+=n ,r = 4 +  3 = 7

}

func External() string {
	fmt.Println(Func220810(3))
	return "success"
}
