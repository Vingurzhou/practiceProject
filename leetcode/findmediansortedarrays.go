/**
 * Created by zhouwenzhe on 2023/3/17
 */

package leetcode

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	num3 := append(nums1, nums2...)
	sort.Ints(num3)
	i := len(num3) / 2
	j := len(num3) % 2
	if j == 1 {
		result = float64(num3[i])
	} else {
		result = float64(num3[i-1]+num3[i]) / 2
	}
	return result
}
