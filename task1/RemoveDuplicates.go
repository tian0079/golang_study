package task1

func RemoveDuplicates(nums []int) (int, []interface{}) {
	length := len(nums)
	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	result := make([]interface{}, length)

	length = len(nums)

	for i = 0; i < len(result); i++ {
		if i < length {
			result[i] = nums[i]
		} else {
			result[i] = "_"
		}
	}

	return length, result
}
