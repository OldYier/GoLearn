package test

import "fmt"

func PrintfYhTriangle() {
	yhTriangle(10)
}
func yhTriangle(n int) {

	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	fmt.Println(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if i == 0 {
				arr[i][j] = 1
			} else if j == 0 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
			}
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Println("")
	}
}
