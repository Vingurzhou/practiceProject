package main

import (
	"fmt"
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	sort.Ints(nums3)
	fmt.Println(nums3)
	if len(nums3)%2 == 1 {
		return float64(nums3[int(math.Floor(float64(len(nums3))/2))])
	} else {
		return float64(nums3[int(float64(len(nums3))/2)]+nums3[int(float64(len(nums3))/2-1)]) / 2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
}
