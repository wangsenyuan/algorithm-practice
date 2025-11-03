package p3732

func maxProduct(nums []int) int64 {
	first := 0
	second := 0
	for _, v := range nums {
		if abs(v) >= first {
			second = first
			first = abs(v)
		} else if abs(v) >= second {
			second = abs(v)
		}
	}

	res := abs(first) * abs(second) * 1e5
	return int64(res)
}

func abs(num int) int {
	return max(num, -num)
}
