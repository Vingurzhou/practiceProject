/**
 * Created by zhouwenzhe on 2023/5/10
 */

package leetcode

import (
	"regexp"
)

func isMatch(s string, p string) bool {
	return regexp.MustCompile(p).FindString(s) == s
}
