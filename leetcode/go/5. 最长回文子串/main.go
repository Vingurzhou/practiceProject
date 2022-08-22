package main

import "fmt"

func longestPalindrome(s string) string {
	for i := range s {
		for j := range s[i:] {
			subString := []rune(s[i : i+j+1])
			for l, r := 0, len(subString); l < len(subString)/2; l, r = l+1, r-1 {
				subString[l], subString[r-1] = subString[r-1], subString[l]
			}
			fmt.Println(subString, subString)
		}
	}
	return ""
}

func main() {
	longestPalindrome("babad")
}
