package p3799

import "slices"

func wordSquares(words []string) [][]string {
	slices.Sort(words)

	var res [][]string

	n := len(words)

	check := func(i, j, u, v int) bool {
		if words[i][0] != words[j][0] {
			return false
		}
		if words[i][3] != words[u][0] {
			return false
		}
		if words[j][3] != words[v][0] {
			return false
		}
		if words[u][3] != words[v][3] {
			return false
		}
		return true
	}

	for i := range n {
		for j := range n {
			for u := range n {
				for v := range n {
					arr := []int{i, j, u, v}
					slices.Sort(arr)
					arr = slices.Compact(arr)
					if len(arr) == 4 && check(i, j, u, v) {
						res = append(res, []string{words[i], words[j], words[u], words[v]})
					}
				}
			}
		}
	}

	return res
}
