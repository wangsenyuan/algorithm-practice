package p3852

import (
	"cmp"
	"slices"
)

type pair struct {
	first  int
	second int
}

func minDistinctFreqPair(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	if len(freq) == 1 {
		return []int{-1, -1}
	}

	var arr []pair

	for k, v := range freq {
		arr = append(arr, pair{k, v})
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	res := []int{arr[0].first}
	for i := 1; i < len(arr); i++ {
		if arr[i].second != arr[0].second {
			res = append(res, arr[i].first)
			return res
		}
	}
	return []int{-1, -1}
}
