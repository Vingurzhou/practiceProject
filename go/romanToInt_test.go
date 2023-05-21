/**
 * Created by zhouwenzhe on 2023/5/16
 */

package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"示例 1",
			args{s: "III"}, 3},
		{"示例 6",
			args{s: "MCMXCIV"}, 1994},
		{"示例 7",
			args{s: "IX"}, 9},
		{"示例 8",
			args{s: "MCDLXXVI"}, 1476},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
