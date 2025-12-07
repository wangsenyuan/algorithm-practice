package p3769

import (
	"cmp"
	"slices"
)

type pair struct {
	first  int
	second int
}

func sortByReflection(nums []int) []int {
	n := len(nums)
	arr := make([]pair, n)
	for i, cur := range nums {
		arr[i] = pair{reverse(cur), cur}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})
	ans := make([]int, n)
	for i := range n {
		ans[i] = arr[i].second
	}
	return ans
}

func reverse(num int) int {
	var res int
	for num > 0 {
		res = res<<1 + num&1
		num >>= 1
	}
	return res
}
