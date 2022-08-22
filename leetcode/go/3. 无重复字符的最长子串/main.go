package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	length := 0
	for i := range s {
		for j := range s[i:] {
			m := make(map[rune]int)
			for _, k := range string(s[i : i+j+1]) {
				if _, ok := m[k]; ok {
					if len(m) > length {
						length = len(m)
					}
					break
				}
				m[k] = 0
			}
			if len(m) == len(s[i:i+j+1]) && len(m) > length {
				length = len(m)
			}
		}
	}
	return length
}

func main() {
	fmt.Println(lengthOfLongestSubstring("otodinokzfhycbuwygqsofctljsgezbvsryceomdvvdyzzuxfnrwstpgejmlkpgegggnuusrswprxmqdzhzrcqzgcltmcz"))
}
