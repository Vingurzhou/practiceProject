package solution

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if result := twoSum([]int{2, 7, 11, 15}, 9); !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("示例 1 错误")
	}
	if result := twoSum([]int{3, 2, 4}, 6); !reflect.DeepEqual(result, []int{1, 2}) {
		t.Errorf("示例 2 错误")
	}
	if result := twoSum([]int{3, 3}, 6); !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("示例 3 错误")
	}

}
