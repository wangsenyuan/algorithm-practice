package p3923

import "slices"

func minGenerations(points [][]int, target []int) int {
	if slices.Max(target) > 6 {
		return -1
	}

	var dim [7][7][7]int

	for _, p := range points {
		dim[p[0]][p[1]][p[2]]++
	}

	for k := 0; ; k++ {
		var cur [][]int
		for i := range 7 {
			for j := range 7 {
				for u := range 7 {
					if dim[i][j][u] > 0 {
						cur = append(cur, []int{i, j, u})
						if i == target[0] && j == target[1] && u == target[2] {
							return k
						}
					}
				}
			}
		}

		var add bool

		for i := range cur {
			for j := range i {
				m0 := (cur[i][0] + cur[j][0]) / 2
				m1 := (cur[i][1] + cur[j][1]) / 2
				m2 := (cur[i][2] + cur[j][2]) / 2
				if dim[m0][m1][m2] == 0 {
					dim[m0][m1][m2]++
					add = true
				}
			}
		}

		if !add {
			break
		}
	}

	return -1
}
