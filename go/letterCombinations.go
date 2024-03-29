/**
 * Created by zhouwenzhe on 2023/6/5
 */

package leetcode

import (
	"strings"
)

func letterCombinations(digits string) []string {
	var ans []string
	hashMap := map[uint8][]uint8{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	if len(digits) == 0 {
		return []string{}

	}
	if len(digits) == 1 {
		for _, character := range hashMap[digits[0]] {
			ans = append(ans, strings.Join([]string{string(character)}, ""))
		}
	}
	subAns := letterCombinations(digits[1:])
	for _, character := range hashMap[digits[0]] {
		for _, sub := range subAns {
			ans = append(ans, strings.Join(append([]string{string(character)}, sub), ""))
		}
	}
	return ans
}
