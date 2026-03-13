package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	costs := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &costs[i])
	}
	races := make([][]int, m)
	for i := range m {
		races[i] = make([]int, 3)
		fmt.Fscan(reader, &races[i][0], &races[i][1], &races[i][2])
	}
	return solve(costs, races)
}

func solve(costs []int, races [][]int) int {
	n := len(costs)

	endAt := make([][]int, n+1)
	for i, cur := range races {
		r := cur[1]
		endAt[r] = append(endAt[r], i)
	}

	tr := NewTree(n + 1)
	var sum int
	var dp int

	for i := 1; i <= n; i++ {
		tr.Update(i, i, sum+dp)
		sum += costs[i-1]

		for _, j := range endAt[i] {
			l, p := races[j][0], races[j][2]
			tr.Update(0, l, p)
		}

		dp = max(dp, tr.Query(0, i)-sum)
	}

	return dp
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &Tree{val, lazy, n}
}

func (t *Tree) apply(i int, v int) {
	t.val[i] += v
	t.lazy[i] += v
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.apply(2*i+1, t.lazy[i])
		t.apply(2*i+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) pull(i int) {
	t.val[i] = max(t.val[2*i+1], t.val[2*i+2])
}

func (t *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if R < l || r < L {
			return
		}
		if L <= l && r <= R {
			t.apply(i, v)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		f(2*i+1, l, mid)
		f(2*i+2, mid+1, r)
		t.pull(i)
	}
	f(0, 0, t.sz-1)
}

func (t *Tree) Query(L int, R int) int {
	var f func(i int, l int, r int) int

	f = func(i int, l int, r int) int {
		if R < l || r < L {
			return 0
		}
		if L <= l && r <= R {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) >> 1
		return max(f(2*i+1, l, mid), f(2*i+2, mid+1, r))
	}
	return f(0, 0, t.sz-1)
}
