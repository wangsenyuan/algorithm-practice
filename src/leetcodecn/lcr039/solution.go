package lcr039

import (
	"cmp"
	"slices"
)

type pair struct {
	first  int
	second int
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	n := len(heights)
	stack := make([]int, n)
	var top int
	var best int
	for i, v := range heights {
		for top > 0 && heights[stack[top-1]] > v {
			h := heights[stack[top-1]]
			l := -1
			if top > 1 {
				l = stack[top-2]
			}
			best = max(best, h*(i-l-1))
			top--
		}

		stack[top] = i
		top++
	}
	return best
}

func largestRectangleArea1(heights []int) int {
	n := len(heights)
	arr := make([]pair, n)
	for i, v := range heights {
		arr[i] = pair{v, i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(b.first-a.first, a.second-b.second)
	})

	L := make([]int, n)
	R := make([]int, n)
	for i := range L {
		L[i] = i
		R[i] = i
	}

	find := func(fa []int, x int) int {
		y := fa[x]
		for y != fa[y] {
			y = fa[y]
		}
		for fa[x] != y {
			x, fa[x] = fa[x], y
		}
		return y
	}

	marked := make([]bool, n)

	var best int
	for _, p := range arr {
		h, i := p.first, p.second
		if i > 0 && marked[i-1] {
			l := find(L, i-1)
			L[i] = l
			R[l] = i
		}
		if i+1 < n && marked[i+1] {
			r := find(R, i+1)
			R[i] = r
			L[r] = i
		}
		l := find(L, i)
		r := find(R, i)
		best = max(best, h*(r-l+1))
		marked[i] = true
	}

	return best
}
