/**
 * @Author: mariomang
 * @Date: 2023-02-24 17:29:39
 * @File: internal/leetcode/leetcode_1_test.go
**/

package leetcode

import (
	"reflect"
	"testing"
)

func TestLeetCode1_TwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "示例2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "示例3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "测试1",
			args: args{
				nums:   []int{2, 5, 5, 11},
				target: 10,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeetCode1_TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LeetCode1_TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
