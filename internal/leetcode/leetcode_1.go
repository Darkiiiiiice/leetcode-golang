/**
 * @Author: mariomang
 * @Date: 2023-02-24 17:25:38
 * @File: internal/leetcode/leetcode_1.go
**/

package leetcode

// LeetCode1_TwoSum 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出
// 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
// PS:
// 1. 2 <= nums.length <= 10^4
// 2. -10^9 <= nums[i] <= 10^9
// 3. -10^9 <= target <= 10^9
// 4. 只会存在一个有效答案
func LeetCode1_TwoSum(nums []int, target int) []int {
	var hash = make(map[int]int)
	for i, num := range nums {
		if n, ok := hash[num]; ok {
			return []int{n, i}
		}
		hash[target-num] = i
	}
	return nil
}
