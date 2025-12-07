package p3771

import "sort"

func totalScore(hp int, damage []int, requirement []int) int64 {
	n := len(damage)
	pref := make([]int, n+1)
	for i, v := range damage {
		pref[i+1] = pref[i] + v
	}

	var res int
	// 看i的贡献
	for i := range n {
		j := sort.Search(i+1, func(j int) bool {
			return hp-(pref[i+1]-pref[j]) >= requirement[i]
		})
		res += i - j + 1
	}

	return int64(res)
}
