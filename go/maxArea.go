/**
 * Created by zhouwenzhe on 2023/5/10
 */

package leetcode

func maxArea(height []int) int {
	var answer int
	leftPointer := 0
	rightPointer := len(height) - 1
	for leftPointer < rightPointer {
		area := min(height[leftPointer], height[rightPointer]) * (rightPointer - leftPointer)
		answer = max(answer, area)
		if height[leftPointer] <= height[rightPointer] {
			leftPointer++
		} else {
			rightPointer--
		}
	}
	return answer
}
