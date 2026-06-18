package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var ac, dr int
	fmt.Fscan(reader, &ac, &dr)
	var n int
	fmt.Fscan(reader, &n)
	a := readNNums(reader, n)
	d := readNNums(reader, n)
	var m int
	fmt.Fscan(reader, &m)
	updates := make([][]int, m)
	for i := range m {
		updates[i] = readNNums(reader, 3)
	}
	return solve(ac, dr, a, d, updates)
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

type pair struct {
	first  int
	second int
}

func solve(ac, dr int, a, d []int, updates [][]int) []int {
	var arr []pair
	n := len(a)

	for i := range n {
		if a[i] > ac || d[i] > dr {
			diff := max(0, a[i]-ac) + max(0, d[i]-dr)
			arr = append(arr, pair{diff, i})
		}
	}

	for i, cur := range updates {
		ai, di := cur[1], cur[2]
		if ai > ac || di > dr {
			diff := max(0, ai-ac) + max(0, di-dr)
			arr = append(arr, pair{diff, i + n})
		}
		// else ignore for now
	}

	slices.SortFunc(arr, func(x pair, y pair) int {
		return cmp.Or(x.first-y.first, x.second-y.second)
	})

	m := len(arr)

	// tr_pos := NewSegTree(m, 0, add)
	tr_min := NewTree(max(m, 1))
	var p0 int
	pos := make([]int, n)

	play := func(i int, p pair) {
		// 不能有重叠的位置
		j := sort.Search(m, func(j int) bool {
			return arr[j].first > p.first || arr[j].first == p.first && arr[j].second >= p.second
		})
		// 这个位置后面的部分,值需要减去1
		tr_min.Update(j+1, m-1, -1)
		tr_min.Set(j, p.first)
		pos[i] = j
	}

	for i := range n {
		pos[i] = -1
		if a[i] > ac || d[i] > dr {
			play(i, pair{max(0, a[i]-ac) + max(0, d[i]-dr), i})
		} else {
			p0++
		}
	}

	cancel := func(i int) {
		if pos[i] < 0 {
			p0--
			return
		}
		// diff := max(0, a[i]-ac) + max(0, d[i]-dr)
		j := pos[i]
		tr_min.Update(j+1, m-1, 1)
		tr_min.Set(j, 0)
		pos[i] = -1
	}

	ans := make([]int, len(updates))

	find := func() int {
		if p0 == 0 {
			return 0
		}
		return tr_min.GetBest(p0)
	}

	for i, cur := range updates {
		id, x, y := cur[0]-1, cur[1], cur[2]
		cancel(id)
		if x > ac || y > dr {
			play(id, pair{max(0, x-ac) + max(0, y-dr), i + n})
		} else {
			p0++
			pos[id] = -1
		}
		// now binary search
		ans[i] = find() + p0
	}

	return ans
}

type Tree struct {
	val  []int
	cnt  []int
	lazy []int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	cnt := make([]int, 4*n)
	lazy := make([]int, 4*n)
	for i := range 4 * n {
		val[i] = 0
		lazy[i] = 0
		cnt[i] = 0
	}
	return &Tree{val, cnt, lazy}
}

func (tr *Tree) apply(i int, v int) {
	tr.val[i] += v
	tr.lazy[i] += v
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.apply(i*2+1, tr.lazy[i])
		tr.apply(i*2+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) Update(L int, R int, v int) {
	if L > R {
		return
	}
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(R, mid))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(L, mid+1), R)
		}
		tr.val[i] = max(tr.val[i*2+1], tr.val[i*2+2])
		tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
	}

	n := len(tr.val) / 4
	f(0, 0, n-1, L, R)
}

func (tr *Tree) Set(pos int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			// L == R
			tr.val[i] = v + tr.lazy[i]
			if v > 0 {
				tr.cnt[i]++
			} else {
				tr.cnt[i]--
			}
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if pos <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.val[i] = max(tr.val[i*2+1], tr.val[i*2+2])
		tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
	}

	n := len(tr.val) / 4
	f(0, 0, n-1)
}

func (tr *Tree) GetBest(p0 int) int {
	var f func(i int, l int, r int) int

	f = func(i int, l int, r int) int {
		if tr.val[i] <= p0 {
			return tr.cnt[i]
		}
		// tr.val[i] > p0
		if l == r {
			return 0
		}
		tr.push(i)
		mid := (l + r) >> 1
		if tr.val[i*2+1] > p0 {
			return f(i*2+1, l, mid)
		}
		return tr.cnt[i*2+1] + f(i*2+2, mid+1, r)
	}

	n := len(tr.val) / 4
	return f(0, 0, n-1)
}
