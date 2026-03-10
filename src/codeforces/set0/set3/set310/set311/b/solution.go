package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m, p int
	fmt.Fscan(reader, &n, &m, &p)

	d := make([]int, n)
	for i := 2; i <= n; i++ {
		fmt.Fscan(reader, &d[i-1])
	}

	h := make([]int, m)
	t := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &h[i], &t[i])
	}

	return solve_cht(p, d, t, h)
}

func solve(p int, d []int, t []int, h []int) int {
	a, pref := prepare(d, t, h)
	m := len(t)
	p = min(p, m)
	prev := make([]int, m+1)
	cur := make([]int, m+1)
	for i := 1; i <= m; i++ {
		prev[i] = inf
	}

	ans := inf
	for groups := 1; groups <= p; groups++ {
		for i := 0; i < groups; i++ {
			cur[i] = inf
		}
		var compute func(left int, right int, optL int, optR int)
		compute = func(left int, right int, optL int, optR int) {
			if left > right {
				return
			}

			mid := (left + right) / 2
			bestPos := -1
			bestVal := inf
			upper := min(mid-1, optR)
			for j := optL; j <= upper; j++ {
				if prev[j] >= inf {
					continue
				}
				cand := prev[j] + (mid-j)*a[mid] - (pref[mid] - pref[j])
				if cand < bestVal {
					bestVal = cand
					bestPos = j
				}
			}

			if groups == 1 {
				bestVal = mid*a[mid] - pref[mid]
				bestPos = 0
			}

			cur[mid] = bestVal
			compute(left, mid-1, optL, bestPos)
			compute(mid+1, right, bestPos, optR)
		}
		compute(groups, m, groups-1, m-1)
		ans = min(ans, cur[m])
		prev, cur = cur, prev
	}

	return ans
}

func solve_cht(p int, d []int, t []int, h []int) int {
	type line struct {
		m int
		b int
	}

	value := func(ln line, x int) int {
		return ln.m*x + ln.b
	}

	bad := func(a line, b line, c line) bool {
		return (c.b-a.b)*(a.m-b.m) <= (b.b-a.b)*(a.m-c.m)
	}

	addLine := func(hull []line, ln line) []line {
		for len(hull) >= 2 && bad(hull[len(hull)-2], hull[len(hull)-1], ln) {
			hull = hull[:len(hull)-1]
		}
		return append(hull, ln)
	}

	advanceHead := func(hull []line, head int, x int) int {
		for head+1 < len(hull) && value(hull[head+1], x) <= value(hull[head], x) {
			head++
		}
		return head
	}

	a, pref := prepare(d, t, h)
	m := len(t)
	p = min(p, m)

	prev := make([]int, m+1)
	cur := make([]int, m+1)
	for i := 1; i <= m; i++ {
		prev[i] = inf
	}

	ans := inf
	for groups := 1; groups <= p; groups++ {
		for i := 0; i < groups; i++ {
			cur[i] = inf
		}
		hull := make([]line, 0, m+1)
		head := 0
		for j := groups - 1; j < m; j++ {
			if prev[j] < inf {
				hull = addLine(hull, line{-j, prev[j] + pref[j]})
			}

			i := j + 1
			if head >= len(hull) {
				cur[i] = inf
				continue
			}
			head = advanceHead(hull, head, a[i])
			best := value(hull[head], a[i])
			cur[i] = i*a[i] - pref[i] + best
		}
		ans = min(ans, cur[m])
		prev, cur = cur, prev
	}

	return ans
}

func prepare(d []int, t []int, h []int) ([]int, []int) {
	n := len(d)
	dist := make([]int, n)
	for i := 1; i < n; i++ {
		dist[i] = dist[i-1] + d[i]
	}

	m := len(t)
	a := make([]int, m+1)
	for i := 1; i <= m; i++ {
		a[i] = t[i-1] - dist[h[i-1]-1]
	}

	sort.Slice(a[1:], func(i, j int) bool {
		return a[i+1] < a[j+1]
	})

	pref := make([]int, m+1)
	for i := 1; i <= m; i++ {
		pref[i] = pref[i-1] + a[i]
	}

	return a, pref
}

const inf = 1 << 60
