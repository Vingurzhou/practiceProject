/**
 * Created by zhouwenzhe on 2023/5/9
 */

package leetcode

import (
	"regexp"
)

func myAtoi(s string) int {
	regex := regexp.MustCompile(`^\s*[+-]?\d+`)
	numStr := regex.FindString(s)
	if numStr == "" {
		return 0
	}
	num := 0
	sign := 1
	for _, character := range numStr {
		switch character {
		case ' ':
		case '+':
		case '-':
			sign = -1
		default:
			num = num*10 + int(character-'0')
			switch {
			case num*sign > 1<<31-1:
				num = 1<<31 - 1
			case num*sign < -1<<31:
				num = 1 << 31
			}
		}
	}
	return num * sign
}
