package p3924

import (
	"math"
	"sort"
)

func minimumThreshold(n int, edges [][]int, source int, target int, k int) int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	maxWt := 0
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
		maxWt = max(maxWt, wt)
	}

	dis := make([]int, n)
	ans := sort.Search(maxWt+1, func(threshold int) bool {
		for i := range dis {
			dis[i] = math.MaxInt
		}
		dis[source] = 0

		type pair struct{ x, d int }
		ql, qr := []pair{{source, dis[source]}}, []pair{} // 模拟双端队列

		for len(ql) > 0 || len(qr) > 0 {
			var p pair
			if len(ql) > 0 {
				ql, p = ql[:len(ql)-1], ql[len(ql)-1] // 队首出
			} else {
				p, qr = qr[0], qr[1:] // 队尾出
			}

			x := p.x
			if x == target {
				return true
			}

			if p.d > dis[x] {
				continue
			}

			for _, e := range g[x] {
				y := e.to
				wt := 0
				if e.wt > threshold {
					wt = 1
				}
				newDis := p.d + wt
				if newDis < dis[y] {
					dis[y] = newDis
					if wt == 0 {
						ql = append(ql, pair{y, newDis}) // 加到队首
					} else if newDis <= k {
						qr = append(qr, pair{y, newDis}) // 加到队尾
					}
				}
			}
		}

		return false
	})

	if ans > maxWt { // 图不连通
		return -1
	}
	return ans
}
