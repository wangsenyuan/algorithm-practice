package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	people := make([][]int, n)
	for i := range n {
		var t, s, f int
		fmt.Fscan(reader, &t, &s, &f)
		people[i] = []int{t, s, f}
	}
	return solve(m, people)
}

type data struct {
	id int
	s  int
	f  int
}

func solve(m int, people [][]int) []int {

	var timestamps []int

	timestamps = append(timestamps, 1)

	for _, p := range people {
		timestamps = append(timestamps, p[0])
	}
	timestamps = append(timestamps, inf)

	sort.Ints(timestamps)
	timestamps = slices.Compact(timestamps)

	n := len(timestamps)
	at := make([][]data, n)
	for i, p := range people {
		t := p[0]
		j := sort.SearchInts(timestamps, t)
		at[j] = append(at[j], data{i, p[1], p[2]})
	}

	ans := make([]int, len(people))
	x := 1

	que := make([][]int, m+1)
	leave := make([][]int, m+1)
	floors := BuildTree(m + 2)
	elevator := BuildTree(m + 2)

	for i, t := range timestamps {
		if i == n-1 {
			break
		}
		// 从t0到时刻t
		for _, p := range at[i] {
			// 在时刻t到达的人,在对应floor排队
			floors.Update(p.s, 1)
			que[p.s] = append(que[p.s], p.id)
		}

		nt := timestamps[i+1]

		for t < nt {
			// 到达x层的人，离梯
			for _, j := range leave[x] {
				ans[j] = t
				elevator.Update(x, -1)
			}
			leave[x] = leave[x][:0]

			// 等待在x层的人，先上梯
			for _, j := range que[x] {
				p := people[j]
				s, f := p[1], p[2]
				floors.Update(s, -1)
				elevator.Update(f, 1)
				leave[f] = append(leave[f], j)
			}
			que[x] = que[x][:0]

			c1, p1 := floors.Query(x+1, m+1)
			c2, p2 := floors.Query(0, x-1)
			c3, p3 := elevator.Query(x+1, m+1)
			c4, p4 := elevator.Query(0, x-1)

			if c1+c3+c2+c4 == 0 {
				// 即无人乘梯，也无人等待
				break
			}

			if c1+c3 >= c2+c4 {
				y := min(p1[0], p3[0])
				dt := min(y-x, nt-t)
				x += dt
				t += dt
			} else {
				y := max(p2[1], p4[1])
				dt := min(x-y, nt-t)
				x -= dt
				t += dt
			}
		}
	}

	return ans
}

const inf = 1 << 60

type Tree struct {
	cnt []int
	pos [][2]int
	n   int
}

func BuildTree(n int) *Tree {
	cnt := make([]int, 4*n)
	pos := make([][2]int, 4*n)
	for i := range 4 * n {
		pos[i][0] = inf
		pos[i][1] = -1
	}
	return &Tree{cnt, pos, n}
}

func (tr *Tree) pull(i int) {
	tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
	tr.pos[i][0] = min(tr.pos[i*2+1][0], tr.pos[i*2+2][0])
	tr.pos[i][1] = max(tr.pos[i*2+1][1], tr.pos[i*2+2][1])
}

func (tr *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.cnt[i] += v
			if tr.cnt[i] == 0 {
				tr.pos[i][0] = inf
				tr.pos[i][1] = -1
			} else {
				tr.pos[i][0] = l
				tr.pos[i][1] = l
			}
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.pull(i)
	}
	f(0, 0, tr.n-1)
}

func (tr *Tree) Query(L int, R int) (cnt int, pos [2]int) {
	var f func(i int, l int, r int, L int, R int) (int, [2]int)
	f = func(i int, l int, r int, L int, R int) (cnt int, pos [2]int) {
		if L == l && R == r {
			return tr.cnt[i], tr.pos[i]
		}
		mid := (l + r) >> 1
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return f(i*2+2, mid+1, r, L, R)
		}
		cnt1, pos1 := f(i*2+1, l, mid, L, mid)
		cnt2, pos2 := f(i*2+2, mid+1, r, mid+1, R)
		cnt = cnt1 + cnt2
		pos[0] = min(pos1[0], pos2[0])
		pos[1] = max(pos1[1], pos2[1])
		return
	}
	return f(0, 0, tr.n-1, L, R)
}
