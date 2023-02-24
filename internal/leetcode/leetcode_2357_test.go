/**
 * @Author: mariomang
 * @Date: 2023-02-24 15:37:49
 * @File: internal/leetcode/leetcode_2357_test.go
**/

package leetcode

import "testing"

func TestLeetCode2357_MinimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				nums: []int{1, 5, 0, 3, 5},
			},
			want: 3,
		},
		{
			name: "示例2",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeetCode2357_MinimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("LeetCode2357_MinimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
