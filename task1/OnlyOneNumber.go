package task1

// 只出现一次的数字
func OnlyOneNumber(nums []int64) int64 {
	var numsMap map[int64]int
	numsMap = make(map[int64]int)
	for _, v := range nums {
		numsMap[v]++
	}

	for num, count := range numsMap {
		if count == 1 {
			return num
		}
	}
	return 0
}
