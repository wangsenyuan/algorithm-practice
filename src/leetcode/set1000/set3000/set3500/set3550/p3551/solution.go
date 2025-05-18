package p3551

import (
	"cmp"
	"slices"
)

func minSwaps(nums []int) int {
	n := len(nums)
	arr := make([]int, n)
	copy(arr, nums)
	slices.SortFunc(arr, func(a, b int) int {
		return cmp.Or(digitSum(a)-digitSum(b), a-b)
	})

	pos := make(map[int]int)
	for i, num := range arr {
		pos[num] = i
	}

	for i := 0; i < n; i++ {
		nums[i] = pos[nums[i]]
	}

	marked := make([]bool, n)
	var res int

	for i := 0; i < n; i++ {
		if !marked[i] {
			j := i
			var cnt int
			for !marked[j] {
				marked[j] = true
				j = nums[j]
				cnt++
			}
			res += cnt - 1
		}
	}

	return res
}

func digitSum(num int) int {
	var res int
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}
