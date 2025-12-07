package p3768

import (
	"slices"
	"sort"
)

func minInversionCount(nums []int, k int) int64 {
	n := len(nums)
	arr := slices.Clone(nums)
	slices.Sort(arr)
	arr = slices.Compact(arr)

	pos := make([]int, n)
	for i := range n {
		pos[i] = sort.SearchInts(arr, nums[i])
	}

	m := len(arr)
	set := make(BIT, m+3)
	var sum int
	for i := range k {
		sum += set.query(pos[i]+1, m)
		set.update(pos[i], 1)
	}

	res := sum

	for i := k; i < n; i++ {
		set.update(pos[i-k], -1)
		sum -= set.query(0, pos[i-k]-1)
		sum += set.query(pos[i]+1, m)
		set.update(pos[i], 1)

		res = min(res, sum)
	}

	return int64(res)
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) query(l int, r int) int {
	return bit.get(r) - bit.get(l-1)
}
