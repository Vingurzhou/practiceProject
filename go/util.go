/**
 * Created by zhouwenzhe on 2023/5/16
 */

package leetcode

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
