package p3767

import "slices"

type pair struct {
	first  int
	second int
}

func maxPoints(technique1 []int, technique2 []int, k int) int64 {
	n := len(technique1)

	var res int
	var cnt int

	var arr2 []pair
	for i := range n {
		if technique1[i] > technique2[i] {
			cnt++
		} else {
			arr2 = append(arr2, pair{technique1[i], technique2[i]})
		}
		res += max(technique1[i], technique2[i])
	}

	if cnt < k {
		// 那么需要交换几个没有被选中的
		slices.SortFunc(arr2, func(a, b pair) int {
			x := a.first - a.second
			y := b.first - b.second
			// 越小越好
			return y - x
		})
		for i := range k - cnt {
			res += arr2[i].first - arr2[i].second
		}
	}

	return int64(res)
}
