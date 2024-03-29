/**
 * Created by zhouwenzhe on 2023/2/3
 */

package leetcode

func twoSum(nums []int, target int) []int {
	lk := -1
	rk := -1
	hashMap := make(map[int]int)
	var result []int
	for k, v := range nums {
		if v2, ok := hashMap[target-v]; ok {
			lk = v2
			rk = k
			result = []int{lk, rk}
			break
		}
		lk = k
		hashMap[v] = lk
	}
	return result
}
