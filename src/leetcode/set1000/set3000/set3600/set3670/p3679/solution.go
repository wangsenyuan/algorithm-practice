package p3679

import "slices"

func minArrivalsToDiscard(arrivals []int, w int, m int) int {
	x := slices.Max(arrivals)
	freq := make([]int, x+1)
	var res int
	added := make([]bool, len(arrivals))
	for i, v := range arrivals {
		if i-w >= 0 && added[i-w] {
			freq[arrivals[i-w]]--
		}
		added[i] = true
		freq[v]++
		if freq[v] > m {
			freq[v]--
			added[i] = false
			res++
		}
	}
	return res
}
