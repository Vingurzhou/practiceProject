package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j, y := range nums[1:] {
			if x+y == target && i != j+1 {
				return []int{i, j + 1}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
