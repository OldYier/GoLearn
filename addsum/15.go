package addsum

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func ThreeSum(nums []int) [][]int {
	var res [][]int
	// 先把数组排序
	sort.Ints(nums)
	fmt.Printf("nums : %v\n", nums)
	var reA []int
	for i := 0; nums[i] < 0; i++ {
		reA = append(reA, nums[i])
	}
	fmt.Printf("reA:%v\n", reA)
	reB := []int{}
	for i := len(nums) - 1; nums[i] >= 0; i-- {
		reB = append(reB, nums[i])
	}
	fmt.Printf("reB %v \n", reB)
	for i := 0; i < len(reA); i++ {
		for j := len(reB) - 1; j >= 0; j-- {
			a := 0
			if i == 2 {
				a++
			}
			if j == 3 {
				a++
			}
			fmt.Printf("i:%v\t", i)
			fmt.Printf("j:%v\t", j)
			fmt.Printf("i+j:%v\n", i+j+a)

		}
	}
	return res
}

// @lc code=end
