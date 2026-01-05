package p3801

func minMergeCost(lists [][]int) int64 {
	// dp[state] 表示ustate表示的位置，合并后的最小cost
	n := len(lists)
	N := 1 << n
	// fp[state][x]表示state组成的序列的计数
	fp := make([][]int, N)

	med := make([]int, N)
	// 先算出中位数吧

	merge := func(s1 int, id int) {
		s := s1 | 1<<id
		if len(fp[s]) > 0 {
			return
		}
		vs := lists[id]
		var res []int
		for i, j := 0, 0; i < len(fp[s1]) || j < len(vs); {
			if j == len(vs) || i < len(fp[s1]) && fp[s1][i] <= vs[j] {
				res = append(res, fp[s1][i])
				i++
			} else {
				res = append(res, vs[j])
				j++
			}
		}

		fp[s] = res
		med[s] = res[(len(res)-1)/2]
	}

	for s := range N {
		for i := range n {
			if (s>>i)&1 == 0 {
				merge(s, i)
			}
		}
	}

	dp := make([]int, N)

	for i := range N {
		dp[i] = inf
	}

	for i := range n {
		dp[1<<i] = 0
	}

	for s := range N {
		for x := s; x > 0; x = (x - 1) & s {
			y := s ^ x
			if y > 0 {
				dp[s] = min(dp[s], dp[x]+dp[y]+len(fp[s])+abs(med[x]-med[y]))
			}
		}
	}

	return int64(dp[N-1])
}

const inf = 1 << 60

func abs(num int) int {
	return max(num, -num)
}
