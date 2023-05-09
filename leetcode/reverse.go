/**
 * Created by zhouwenzhe on 2023/5/9
 */

package leetcode

import (
	"math"
)

func reverse(x int) int {
	var result int
	for {
		quotient := x / 10
		remainder := x % 10
		x = quotient
		result = result*10 + remainder
		if quotient == 0 {
			break
		}
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
