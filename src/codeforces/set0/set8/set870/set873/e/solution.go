package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (a []int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a)
	return
}

const inf = 1 << 60

type student struct {
	id    int
	score int
}

// d1 for d1 - c2
// d2 for d2 - c3
// d3 for d3 - c(-1)
type data struct {
	d [3]int
}

func solve(a []int) []int {
	n := len(a)
	students := make([]student, n)
	for i := range n {
		students[i] = student{id: i, score: a[i]}
	}
	slices.SortFunc(students, func(x student, y student) int {
		return cmp.Or(y.score-x.score, x.id-y.id)
	})

	dp := make([]SegmentTree, n)

	for i := n - 1; i >= 0; i-- {
		dp[i] = NewSegTree(n)
		for j := i; j < n; j++ {
			// 如果i..j rank 3
			v := students[j].score
			if j+1 < n {
				v -= students[j+1].score
			}
			dp[i].Update(j-i, v)
		}
	}

	// d1 is fixed
	arr := make([]int, 3)
	d := []int{-inf, -inf, -inf}

	play := func(f int) bool {
		found := false
		for i := f; i < n-1; i++ {
			w := max(f, i-f+1)
			v := min(f, i-f+1)
			if w > 2*v {
				continue
			}
			// v <= 2 * w must hold
			l := (w + 1) / 2
			if l > n-i-1 {
				// 不够rank 3
				break
			}
			r := min(v*2, n-i-1)
			if d[1] <= students[i].score-students[i+1].score {
				if d[1] < students[i].score-students[i+1].score {
					// reset
					d[2] = -inf
				}
				d2 := dp[i+1].Get(l-1, r)
				if d[2] <= d2.first {
					d[1] = students[i].score - students[i+1].score
					d[2] = d2.first
					arr[1] = i - f + 1
					arr[2] = d2.second + 1
					found = true
					// 还必须知道位置
				}
			}
		}
		return found
	}

	for i := range n - 1 {
		w := i + 1
		v := (w + 1) / 2 * 2
		if n-w >= v && d[0] <= students[i].score-students[i+1].score {
			if d[0] < students[i].score-students[i+1].score {
				// reset
				d[1] = -inf
				d[2] = -inf
			}
			if play(w) {
				d[0] = students[i].score - students[i+1].score
				arr[0] = i + 1
			}
		}
	}

	res := make([]int, n)

	for i := range n {
		res[students[i].id] = -1
		if i < arr[0] {
			res[students[i].id] = 1
		} else if i < arr[0]+arr[1] {
			res[students[i].id] = 2
		} else if i < arr[0]+arr[1]+arr[2] {
			res[students[i].id] = 3
		}
	}

	return res
}

type pair struct {
	first  int
	second int
}

func max_pair(a, b pair) pair {
	if a.first > b.first || a.first == b.first && a.second > b.second {
		return a
	}
	return b
}

type SegmentTree []pair

func NewSegTree(n int) SegmentTree {
	res := make(SegmentTree, 2*n)
	for i := n; i < 2*n; i++ {
		res[i] = pair{-inf, i - n}
	}
	for i := n - 1; i > 0; i-- {
		res[i] = max_pair(res[i*2], res[i*2+1])
	}
	return res
}

func (this SegmentTree) Update(p int, v int) {
	p += len(this) >> 1
	this[p].first = v
	for p > 1 {
		this[p>>1] = max_pair(this[p], this[p^1])
		p >>= 1
	}
}

func (this SegmentTree) Get(l int, r int) pair {
	res := pair{-inf, -1}
	l += len(this) >> 1
	r += len(this) >> 1
	for l < r {
		if l&1 == 1 {
			res = max_pair(res, this[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max_pair(res, this[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
