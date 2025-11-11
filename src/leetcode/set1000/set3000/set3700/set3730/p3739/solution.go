package p3739

import (
	"slices"
	"sort"
)

func countMajoritySubarrays(nums []int, target int) int64 {
	var arr []int
	arr = append(arr, 0)
	var sum int
	for i, num := range nums {
		if num == target {
			sum++
		}
		arr = append(arr, 2*sum-(i+1))
	}

	// 2 * (sumr - suml) > (r - l)
	slices.Sort(arr)
	arr = slices.Compact(arr)

	n := len(arr)

	var res int
	bit := make(BIT, n+2)
	bit.update(sort.SearchInts(arr, 0), 1)

	sum = 0
	for i, num := range nums {
		if num == target {
			sum++
		}
		cur := 2*sum - (i + 1)
		j := sort.SearchInts(arr, cur)
		if j > 0 {
			res += bit.rangeQuery(0, j-1)
		}
		bit.update(j, 1)
	}

	return int64(res)
}

type BIT []int

func (bit BIT) update(i int, val int) {
	i++
	for i < len(bit) {
		bit[i] += val
		i += i & -i
	}
}

func (bit BIT) query(i int) int {
	res := 0
	i++
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) rangeQuery(l int, r int) int {
	return bit.query(r) - bit.query(l-1)
}
