/**
 * Created by zhouwenzhe on 2023/3/17

5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。
如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
*/

package leetcode

func longestPalindrome(s string) string {
	result := s[0:1]
	dp := make([][]bool, len(s))
	for k, _ := range dp {
		dp[k] = make([]bool, len(s))
		dp[k][k] = true
	}

	for l := 2; l <= len(s); l++ {
		for i := 0; i <= len(s)-1; i++ {
			j := l + i - 1
			if j >= len(s) {
				break
			}
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}
