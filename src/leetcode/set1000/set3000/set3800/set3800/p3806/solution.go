package p3806

import (
	"slices"
)

const H = 31

func maximumAND(nums []int, k int, m int) int {
	if len(nums) == 1 {
		return nums[0] + k
	}

	var f func(i int, x int) int

	n := len(nums)

	arr := make([]int, n)

	get := func(num int, x int) int {
		for i := H - 1; i >= 0; i-- {
			if (x>>i)&1 == 1 {
				if (num>>i)&1 == 0 {
					return x - num
				}
				// else continue
			} else {
				// x[i] = 0
				if (num>>i)&1 == 1 {
					num ^= 1 << i
				}
			}
		}
		// 两个相同
		return 0
	}

	f = func(i int, x int) int {
		if i < 0 {
			return x
		}

		y := x | 1<<i

		for i, num := range nums {
			arr[i] = get(num, y)
		}

		slices.Sort(arr)

		var sum int
		for i := range m {
			sum += arr[i]
		}
		if sum <= k {
			return f(i-1, y)
		}
		return f(i-1, x)
	}

	return f(H-1, 0)
}
