package addsum

import "fmt"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func ThreeSum(nums []int) [][]int {
	var res [][]int
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			for k := len(nums); k > 1; k-- {
				if nums[i]+nums[j]+nums[k] == 0 {
					var re []int
					re = append(re, nums[i])
					re = append(re, nums[j])
					re = append(re, nums[k])
					// fmt.Printf("re:%v", re)
					res = append(res, re)
				}
			}
		}

	}
	fmt.Printf("res:%v", res)
	return res
}

// @lc code=end
