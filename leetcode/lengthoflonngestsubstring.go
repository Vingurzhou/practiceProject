/**
 * Created by zhouwenzhe on 2023/3/10
 */

package leetcode

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	set := make(map[string]int)
	var maxLength int
	lk := -1
	rk := 0
	for k, _ := range s {
		lk = k
		for rk < len(s) {
			_, ok := set[fmt.Sprintf("%c ", s[rk])]
			if ok {
				if len(set) > maxLength {
					maxLength = len(set)
				}
				break
			}
			set[fmt.Sprintf("%c ", s[rk])] = 0
			rk++
			if len(set) > maxLength {
				maxLength = len(set)
			}
		}
		delete(set, fmt.Sprintf("%c ", s[lk]))
	}
	fmt.Println(maxLength)
	return maxLength
}
