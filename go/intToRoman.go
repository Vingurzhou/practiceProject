/**
 * Created by zhouwenzhe on 2023/5/16
 */

package leetcode

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	var answer string
	for num > 0 {
		for _, i := range valueSymbols {
			if i.value <= num {
				answer += i.symbol
				num -= i.value
				break
			}
		}
	}
	return answer
}
