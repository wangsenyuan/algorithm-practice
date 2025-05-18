package p3550

func smallestIndex(nums []int) int {
	for i, num := range nums {
		if digitSum(num) == i {
			return i
		}
	}
	return -1
}

func digitSum(num int) int {
	var res int
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}
