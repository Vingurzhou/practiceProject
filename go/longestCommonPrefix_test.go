/**
 * Created by zhouwenzhe on 2023/5/16
 */

package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"示例 1",
			args{strs: []string{"flower", "flow", "flight"}},
			"fl"},
		{"示例 3",
			args{strs: []string{""}},
			""},
		{"示例 4",
			args{strs: []string{"flower", "flower", "flower", "flower"}},
			"flower"},
		{"示例 5",
			args{strs: []string{"a", "b"}},
			""},
		{"示例 6",
			args{strs: []string{"ab", "a"}},
			"a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
