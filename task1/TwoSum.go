package task1

func TwoSum(nums []int, target int) ([2]int, bool) {
	flag := false
	var array [2]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(nums[i]+nums[j]) == target {
				array[0] = i
				array[1] = j
				flag = true
				return array, true
			}
		}
	}
	return array, flag
}
