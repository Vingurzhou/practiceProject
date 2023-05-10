/**
 * Created by zhouwenzhe on 2023/5/10
 */

package leetcode

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"示例 1",
			args{
				s: "aa",
				p: "a",
			},
			false},
		{"示例 2",
			args{
				s: "aa",
				p: "a*",
			},
			true},
		{"示例 3",
			args{
				s: "ab",
				p: ".*",
			},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
