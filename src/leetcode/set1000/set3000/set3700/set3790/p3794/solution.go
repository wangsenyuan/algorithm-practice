package p3794

const inf = 1 << 60

func findMaxVal(n int, restrictions [][]int, diff []int) int {
	h := make([]int, n)
	for i := range n {
		h[i] = inf
	}

	for _, cur := range restrictions {
		id, x := cur[0], cur[1]
		h[id] = min(h[id], x)
	}

	h[0] = 0

	for i := range n - 1 {
		h[i+1] = min(h[i+1], h[i]+diff[i])
	}
	res := h[n-1]
	next := h[n-1]
	for i := n - 2; i >= 0; i-- {
		cur := min(h[i], next+diff[i])
		res = max(res, cur)
		next = cur
	}

	return res
}
