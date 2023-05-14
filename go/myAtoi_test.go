/**
 * Created by zhouwenzhe on 2023/5/9
 */

package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"示例 1",
			args{s: "42"},
			42},
		{"示例 2",
			args{s: "   -42"},
			-42},
		{"示例 3",
			args{s: "4193 with words"},
			4193},
		{"示例 4",
			args{s: "words and 987"},
			0},
		{"示例 5",
			args{s: "-91283472332"},
			-2147483648},
		{"示例 6",
			args{s: "-+12"},
			0},
		{"示例 7",
			args{s: "00000-42a1234"},
			0},
		{"示例 8",
			args{s: "9223372036854775808"},
			2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
