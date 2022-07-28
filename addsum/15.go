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
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > 0; j-- {

		}
	}
	return res
}

// @lc code=end
