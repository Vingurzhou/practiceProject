/**
 * Created by zhouwenzhe on 2023/5/16
 */

package leetcode

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	var answer int
	for i := range s {
		if i < len(s)-1 && symbolValues[s[i]] < symbolValues[s[i+1]] {
			answer -= symbolValues[s[i]]
		} else {
			answer += symbolValues[s[i]]
		}
	}
	return answer
}
