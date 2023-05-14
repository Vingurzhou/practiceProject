/**
 * Created by zhouwenzhe on 2023/5/10
 */

package leetcode

import (
	"math"
)

func maxArea(height []int) int {
	var ans int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			ans = int(math.Max(float64(ans), math.Min(float64(height[i]), float64(height[j]))*float64(j-i)))
		}
	}
	return ans
}
