package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, v := range res {
		fmt.Fprintf(writer, "%.6f\n", v)
	}
}

func drive(reader *bufio.Reader) []float64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	tubes := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &tubes[i])
	}
	actions := make([][]int, m)
	for i := range m {
		actions[i] = make([]int, 3)
		fmt.Fscan(reader, &actions[i][0], &actions[i][1])
		if actions[i][0] == 1 {
			fmt.Fscan(reader, &actions[i][2])
		}
	}
	return solve(tubes, actions)
}

func solve(tubes []int, actions [][]int) []float64 {
	// n := len(tubes)

	arr := slices.Clone(tubes)
	for _, cur := range actions {
		if cur[0] == 1 {
			arr = append(arr, cur[2])
		}
	}
	slices.Sort(arr)
	arr = slices.Compact(arr)

	tr := NewTree(arr)
	for _, v := range tubes {
		i := sort.SearchInts(arr, v)
		tr.Update(i, 1)
	}

	update := func(p int, v int) {
		i := sort.SearchInts(arr, tubes[p])
		tr.Update(i, -1)
		tubes[p] = v
		i = sort.SearchInts(arr, v)
		tr.Update(i, 1)
	}

	var ans []float64

	for _, cur := range actions {
		if cur[0] == 1 {
			update(cur[1]-1, cur[2])
		} else {
			v := cur[1]
			f := tr.Find(v)
			ans = append(ans, f)
		}
	}

	return ans
}

type Tree struct {
	a   []int
	sum []int
	cnt []int
	val []int
	sz  int
}

func NewTree(a []int) *Tree {
	n := len(a)
	sum := make([]int, 4*n)
	cnt := make([]int, 4*n)
	val := make([]int, 4*n)
	var dfs func(i int, l int, r int)
	dfs = func(i int, l int, r int) {
		if l == r {
			val[i] = a[l]
			cnt[i] = 0
			sum[i] = 0
			return
		}
		mid := (l + r) >> 1
		dfs(i*2+1, l, mid)
		dfs(i*2+2, mid+1, r)
		// val[i] = val[i*2+2]
	}
	dfs(0, 0, n-1)
	return &Tree{a: a, sum: sum, cnt: cnt, val: val, sz: n}
}

func (tr *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.cnt[i] += v
			tr.sum[i] += v * tr.a[l]
			if tr.cnt[i] == 0 {
				tr.val[i] = 0
			} else {
				tr.val[i] = tr.a[l]
			}
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
		tr.sum[i] = tr.sum[i*2+1] + tr.sum[i*2+2]
		tr.val[i] = max(tr.val[i*2+1], tr.val[i*2+2])
	}
	f(0, 0, tr.sz-1)
}

func (tr *Tree) Find(v int) float64 {
	var f func(i int, l int, r int, cnt int, sum int) float64
	f = func(i int, l int, r int, cnt int, sum int) float64 {
		if l == r {
			need := (cnt+tr.cnt[i])*tr.val[i] - (sum + tr.sum[i])
			w := v - need
			if cnt+tr.cnt[i] == 0 || w < 0 {
				return float64(tr.val[i] + v)
			}
			return float64(w)/float64(cnt+tr.cnt[i]) + float64(tr.val[i])
		}
		mid := (l + r) >> 1
		need := (cnt+tr.cnt[2*i+1])*tr.val[2*i+1] - (sum + tr.sum[i*2+1])
		if v < need {
			return f(i*2+1, l, mid, cnt, sum)
		}
		w := v - need
		// val[i]是对应区间的最大值
		// w > 0
		var res float64 = 1 << 60
		if cnt+tr.cnt[2*i+1] > 0 {
			res = float64(w)/float64(cnt+tr.cnt[2*i+1]) + float64(tr.val[2*i+1])
		}
		return min(res, f(i*2+2, mid+1, r, cnt+tr.cnt[i*2+1], sum+tr.sum[i*2+1]))
	}

	return f(0, 0, tr.sz-1, 0, 0)
}
