package leetcode

import "sort"

// LeetCode2357_MinimumOperations 使数组中所有元素都等于零
// 给你一个非负整数数组 nums 。在一步操作中，你必须：
// * 选出一个正整数 x ，x 需要小于或等于 nums 中 最小 的 非零 元素。
// * nums 中的每个正整数都减去 x。
// 返回使 nums 中所有元素都等于 0 需要的 最少 操作数
// PS:
// 1. 1 <= nums.length <= 100
// 2. 0 <= nums[i] <= 100
func LeetCode2357_MinimumOperations(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var set = make(map[int]struct{})
	for _, n := range nums {
		if n > 0 {
			set[n] = struct{}{}
		}
	}
	return len(set)
}

// LeetCode2357_MinimumOperations_Array 数组法
func LeetCode2357_MinimumOperations_Array(nums []int) int {
	ans := 0
	sort.Ints(nums)
	a := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > a {
			ans++
			a = nums[i]
		}
	}
	return ans
}
