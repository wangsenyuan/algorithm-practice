package p3800

func minimumCost(s string, t string, flipCost int, swapCost int, crossCost int) int64 {
	// 这类题目好麻烦
	// 一次swap可以改变两个位置
	// 比如 00  11 那么一次cross + swap 就可以变成 10 10
	arr := make([][]int, 2)
	for i := range len(s) {
		if s[i] != t[i] {
			d := int(s[i] - '0')
			arr[d] = append(arr[d], i)
		}
	}
	var res int
	// 如果 01 10 可以用swap
	if swapCost <= 2*flipCost {
		// 01 vs 10
		for len(arr[0]) >= 1 && len(arr[1]) >= 1 {
			res += swapCost
			arr[0] = arr[0][1:]
			arr[1] = arr[1][1:]
		}
	}

	if swapCost+crossCost <= 2*flipCost {
		for len(arr[0]) >= 2 {
			res += swapCost + crossCost
			arr[0] = arr[0][2:]
		}
		for len(arr[1]) >= 2 {
			res += swapCost + crossCost
			arr[1] = arr[1][2:]
		}
	}

	res += flipCost*len(arr[0]) + flipCost*len(arr[1])

	return int64(res)
}
