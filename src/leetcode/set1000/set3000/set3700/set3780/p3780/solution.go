package p3780

import "slices"

func maximumSum(nums []int) int {
	arr := make([][]int, 3)
	for _, num := range nums {
		arr[num%3] = append(arr[num%3], num)
	}
	for i := range 3 {
		slices.Sort(arr[i])
		slices.Reverse(arr[i])
	}

	var res int
	for i := range 3 {
		if len(arr[i]) >= 3 {
			res = max(res, arr[i][0]+arr[i][1]+arr[i][2])
		}
	}
	// 0 + 1 + 2
	if len(arr[0]) > 0 && len(arr[1]) > 0 && len(arr[2]) > 0 {
		res = max(res, arr[0][0]+arr[1][0]+arr[2][0])
	}
	return res
}
