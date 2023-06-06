/**
 * Created by zhouwenzhe on 2023/6/5
 */

package leetcode

func threeSumClosest(nums []int, target int) int {
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				temp := nums[i] + nums[j] + nums[k]
				if abs(temp-target) < abs(ans-target) {
					ans = temp
				}
			}
		}
	}
	return ans
}
