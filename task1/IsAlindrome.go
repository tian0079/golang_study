package task1

import (
	"fmt"
)

func IsAlindrome(ints int) bool {
	str := fmt.Sprintf("%d", ints)
	length := len(str)
	if length <= 1 {
		return false
	}
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}
