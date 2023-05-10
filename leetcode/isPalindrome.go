/**
 * Created by zhouwenzhe on 2023/5/10
 */

package leetcode

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		quotient := x / 10
		remainder := x % 10
		x = quotient
		revertedNumber = revertedNumber*10 + remainder
	}
	return x == revertedNumber || x == revertedNumber/10
}
