package EveryDay

import (
	"sync"
)

const N = 10

var wg = &sync.WaitGroup{}

func Func230214() {

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			println(i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
}
