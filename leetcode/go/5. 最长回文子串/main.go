package main

import "fmt"

func longestPalindrome(s string) string {
	length := 0
	var result string
	for i := range s {
		for j := range s[i:] {
			subString := []rune(s[i : i+j+1])
			subString2 := make([]rune, len(subString))
			copy(subString2, subString)
			for l, r := 0, len(subString); l < len(subString)/2; l, r = l+1, r-1 {
				subString[l], subString[r-1] = subString[r-1], subString[l]
			}
			if string(subString) == string(subString2) {
				if len(subString) > length {
					length = len(subString)
					result = string(subString)
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
