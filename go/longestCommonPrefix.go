/**
 * Created by zhouwenzhe on 2023/5/16
 */

package leetcode

func longestCommonPrefix(strs []string) (answer string) {
	defer func() {
		recover()
	}()
	pointer := 0

	for {
		num := 0
		character := strs[0][pointer]
		for i := 0; i < len(strs) && strs[i][pointer] == character; i++ {
			num += 1
		}
		if num != len(strs) {
			break
		}
		answer += string(character)
		pointer += 1
	}
	return answer
}
