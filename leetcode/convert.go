/**
 * Created by zhouwenzhe on 2023/5/8
 */

package leetcode

func convert(s string, numRows int) string {
	t := numRows + numRows - 2
	length := len(s)
	if numRows == 1 || numRows >= length {
		return s
	}
	numColumns := length / t * (numRows - 2 + 1)
	if length%t <= numRows {
		numColumns += 1
	} else {
		numColumns += 1 + length%t - numRows
	}

	arr := make([][]string, numRows)
	for i := range arr {
		arr[i] = make([]string, numColumns)
	}
	x, y := 0, 0
	for index, letter := range s {
		arr[x][y] = string(letter)
		if index%t < numRows-1 {
			x++
		} else {
			x--
			y++
		}
	}
	var result string
	for _, row := range arr {
		for _, letter := range row {
			result += letter
		}
	}
	return result
}
