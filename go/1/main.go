package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for k1, v1 := range nums {
		for k2, v2 := range nums[k1+1:] {
			if v1+v2 == target {
				return []int{k1, k1 + k2 + 1}
			}
		}
	}
	return nil
}

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}
