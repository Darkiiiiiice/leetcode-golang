/**
 * @Author: mariomang
 * @Date: 2023-02-24 18:11:21
 * @File: internal/leetcode/leetcode_4.go
**/

package leetcode

import "sort"

// LeetCode4_FindMedianSortedArrays 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 算法的时间复杂度应该为 O(log (m+n)) 。
// PS:
// 1. nums1.length == m
// 2. nums2.length == n
// 3. 0 <= m <= 1000
// 4. 0 <= n <= 1000
// 5. 1 <= m + n <= 2000
// 6. -10^6 <= nums1[i], nums2[i] <= 10^6
func LeetCode4_FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var arr = make([]int, 0)
	arr = append(arr, nums1...)
	arr = append(arr, nums2...)

	sort.Ints(arr)

	var l = len(arr)
	if len(arr)%2 == 1 {
		return float64(arr[l/2])
	}

	return (float64(arr[l/2]) + float64(arr[(l/2)-1])) / 2
}
