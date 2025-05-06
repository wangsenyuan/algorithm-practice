package p1128

func numEquivDominoPairs(dominoes [][]int) int {
	n := len(dominoes)

	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		tmp := dominoes[i]
		a, b := tmp[0], tmp[1]
		if a > b {
			a, b = b, a
		}
		cnt[a*n+b]++
	}

	var res int

	for _, m := range cnt {
		res += m * (m - 1) / 2
	}

	return res
}
