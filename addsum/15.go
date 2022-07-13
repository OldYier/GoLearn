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
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(reA); j++ {
			fmt.Printf(i+)
		}
	}
	return res
}

// @lc code=end
