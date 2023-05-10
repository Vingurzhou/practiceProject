/**
 * Created by zhouwenzhe on 2023/5/10
 */

package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"示例 1",
			args{x: 121},
			true},
		{"示例 2",
			args{x: -121},
			false},
		{"示例 3",
			args{x: 10},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
