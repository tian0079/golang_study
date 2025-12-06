package task1

func PlusOne(digits []int) []int {

	sum := 0
	//将数组转换成数值
	for i := 0; i < len(digits); i++ {
		sum = sum*10 + digits[i]
	}
	sum += 1

	slice := make([]int, 0)
	for sum > 0 {
		temp := sum % 10 //取出最后一位
		slice = append(slice, temp)
		sum = sum / 10 //去掉最后一位
	}
	//反序切片
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
