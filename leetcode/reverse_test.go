/**
 * Created by zhouwenzhe on 2023/5/9
 */

package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"示例 1",
			args{x: 123},
			321},
		{"示例 2",
			args{x: -123},
			-321},
		{"示例 3",
			args{x: 120},
			21},
		{"示例 4",
			args{x: 0},
			0},
		{"示例 5",
			args{x: 1534236469},
			0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
