/**
 * Created by zhouwenzhe on 2023/5/10
 */

package leetcode

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"示例 1", args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"示例 2", args{height: []int{1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
