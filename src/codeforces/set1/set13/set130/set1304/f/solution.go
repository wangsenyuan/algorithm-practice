package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a, k)
}

func solve(a [][]int, k int) int {
	n := len(a)
	m := len(a[0])
	f := make([]int, m-k+1)
	q := []int{}
	s := make([]int, m+1)
	for i := range n {
		for j := 1; j <= m; j++ {
			s[j] = s[j-1] + a[i][j-1]
		}

		if i == 0 {
			for j := range f {
				f[j] = s[j+k] - s[j]
			}
			continue
		}

		nf := make([]int, len(f))
		mx := 0
		q = q[:0]
		for j, fj := range f {
			if j >= k {
				mx = max(mx, f[j-k]+s[j]-s[j-k])
			}

			for len(q) > 0 && f[q[len(q)-1]]-s[q[len(q)-1]] <= fj-s[j] {
				q = q[:len(q)-1]
			}
			q = append(q, j)
			if q[0] <= j-k {
				q = q[1:]
			}

			nf[j] = max(mx-s[j], f[q[0]]-s[q[0]]) + s[j+k]
		}

		mx = 0
		q = q[:0]
		for j := m - k; j >= 0; j-- {
			if j <= m-k*2 {
				mx = max(mx, f[j+k]+s[j+k*2]-s[j+k])
			}

			for len(q) > 0 && f[q[len(q)-1]]+s[q[len(q)-1]+k] <= f[j]+s[j+k] {
				q = q[:len(q)-1]
			}
			q = append(q, j)
			if q[0] >= j+k {
				q = q[1:]
			}

			nf[j] = max(nf[j], max(mx+s[j+k], f[q[0]]+s[q[0]+k])-s[j])
		}

		f = nf
	}

	return slices.Max(f)
}

func solveTr(a [][]int, k int) int {
	n := len(a)
	if n == 1 {
		return solve1(a[0], k)
	}
	m := len(a[0])

	if k == m {
		return solve2(a)
	}

	tr := NewTree(m)
	var sum int
	for j := 0; j < m; j++ {
		if j >= k {
			sum -= a[0][j-k]
			sum -= a[1][j-k]
		}
		sum += a[0][j]
		sum += a[1][j]
		if j >= k-1 {
			tr.Update(j, j, sum)
		}
	}
	a = append(a, make([]int, m))
	dp := make([]int, m)
	for i := 1; i < n; i++ {
		// i和i+1行的值
		sum = 0
		for j := 0; j < k; j++ {
			sum += a[i][j]
			sum += a[i+1][j]
			tr.Update(k-1, min(m-1, j+k-1), -a[i][j])
		}
		clear(dp)
		dp[k-1] = tr.val[0] + sum

		for j := k; j < m; j++ {
			sum += a[i][j]
			sum += a[i+1][j]
			sum -= a[i][j-k]
			sum -= a[i+1][j-k]
			tr.Update(j, min(m-1, j+k-1), -a[i][j])
			tr.Update(j-k, j-1, a[i][j-k])
			dp[j] = tr.val[0] + sum
		}

		tr.Reset(dp)
	}

	return tr.val[0]
}

func solve2(a [][]int) int {
	var res int
	for _, row := range a {
		for _, v := range row {
			res += v
		}
	}
	return res
}

func solve1(a []int, k int) int {
	var best int

	var sum int
	for i := 0; i < len(a); i++ {
		sum += a[i]
		if i >= k {
			sum -= a[i-k]
		}
		if i >= k-1 {
			best = max(best, sum)
		}
	}

	return best
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	arr := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &Tree{arr, lazy, n}
}

func (t *Tree) update(i int, v int) {
	t.val[i] += v
	t.lazy[i] += v
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.update(i*2+1, t.lazy[i])
		t.update(i*2+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) pull(i int) {
	t.val[i] = max(t.val[i*2+1], t.val[i*2+2])
}

func (t *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.update(i, v)
			return
		}
		t.push(i)
		mid := (l + r) / 2
		if R <= mid {
			f(2*i+1, l, mid, L, R)
		} else if mid < L {
			f(2*i+2, mid+1, r, L, R)
		} else {
			f(2*i+1, l, mid, L, mid)
			f(2*i+2, mid+1, r, mid+1, R)
		}
		t.pull(i)
	}
	f(0, 0, t.sz-1, L, R)
}

func (t *Tree) Reset(arr []int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		t.lazy[i] = 0
		if l == r {
			t.val[i] = arr[l]
			return
		}
		mid := (l + r) / 2
		f(i*2+1, l, mid)
		f(i*2+2, mid+1, r)
		t.pull(i)
	}
	f(0, 0, t.sz-1)
}
