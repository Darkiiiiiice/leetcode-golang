/**
 * @Author: mariomang
 * @Date: 2023-02-24 18:13:25
 * @File: internal/leetcode/leetcode_4_test.go
**/

package leetcode

import "testing"

func TestLeetCode4_FindMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "示例1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeetCode4_FindMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("LeetCode4_FindMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
