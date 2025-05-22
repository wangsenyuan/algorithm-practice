package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	segs := make([][]int, n)
	for i := range n {
		segs[i] = readNNums(reader, 2)
	}
	queries := make([][]int, m)
	for i := range m {
		var cnt int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &cnt) + 1
		queries[i] = make([]int, cnt)
		for j := range cnt {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	return solve(segs, queries)
}

const MAX_R = 1000000

type BIT []int

func (bit BIT) add(p int, v int) {
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) query(p int) int {
	res := 0
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

func (bit BIT) queryRange(l int, r int) int {
	return bit.query(r) - bit.query(l-1)
}

type pair struct {
	first  int
	second int
}

func solve(segs [][]int, queries [][]int) []int {
	xs := make([][]int, MAX_R+1)
	for _, seg := range segs {
		l, r := seg[0], seg[1]
		xs[r] = append(xs[r], l)
	}

	qs := make([][]pair, MAX_R+1)
	for i, cur := range queries {
		var p int
		for _, j := range cur {
			qs[j-1] = append(qs[j-1], pair{p + 1, i})
			p = j
		}
		qs[MAX_R] = append(qs[MAX_R], pair{p + 1, i})
	}

	bit := make(BIT, MAX_R+10)

	ans := make([]int, len(queries))

	for i := 1; i <= MAX_R; i++ {
		for _, l := range xs[i] {
			bit.add(l, 1)
		}
		for _, cur := range qs[i] {
			id := cur.second
			ans[id] -= bit.queryRange(cur.first, i)
		}
	}

	for i := range ans {
		ans[i] += len(segs)
	}
	return ans
}

func solve1(segs [][]int, queries [][]int) []int {
	var nums []int
	for _, seg := range segs {
		nums = append(nums, seg...)
	}

	for _, cur := range queries {
		nums = append(nums, cur...)
	}
	nums = sortAndUnique(nums)

	n := len(nums)

	at := make([][]int, n)
	for _, seg := range segs {
		l, r := seg[0], seg[1]
		r = sort.SearchInts(nums, r)
		l = sort.SearchInts(nums, l)
		at[r] = append(at[r], l)
	}
	roots := make([]*Tree, n+1)

	for i := n - 1; i >= 0; i-- {
		// 复制一个出来
		roots[i] = roots[i+1].Clone()
		for _, l := range at[i] {
			roots[i] = roots[i].Add(0, n-1, l, 1)
		}
	}

	res := make([]int, len(queries))

	for i, cur := range queries {
		p := 1
		for _, j := range cur {
			u := sort.SearchInts(nums, p)
			v := sort.SearchInts(nums, j)
			tmp := roots[v].Query(u, v, 0, n-1)
			res[i] += tmp
			p = j + 1
		}
	}

	return res
}

func sortAndUnique(nums []int) []int {
	sort.Ints(nums)
	var n int
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] > nums[i-1] {
			nums[n] = nums[i-1]
			n++
		}
	}
	return nums[:n]
}

type Tree struct {
	left, right *Tree
	sum         int
}

func (t *Tree) GetSum() int {
	if t == nil {
		return 0
	}
	return t.sum
}

func (t *Tree) Clone() *Tree {
	res := new(Tree)
	if t != nil {
		res.left = t.left
		res.right = t.right
		res.sum = t.sum
	}
	return res
}

func (t *Tree) pull() {
	t.sum = t.left.GetSum() + t.right.GetSum()
}

func (t *Tree) Add(l int, r int, p int, v int) *Tree {
	res := t.Clone()
	if l == r {
		res.sum += v
		return res
	}
	mid := (l + r) / 2
	if p <= mid {
		res.left = res.left.Add(l, mid, p, v)
	} else {
		res.right = res.right.Add(mid+1, r, p, v)
	}
	res.pull()
	return res
}

func (t *Tree) Query(L int, R int, l int, r int) int {
	if t == nil {
		return 0
	}
	if L == l && R == r {
		return t.sum
	}
	mid := (l + r) / 2
	if R <= mid {
		return t.left.Query(L, R, l, mid)
	}
	if mid < L {
		return t.right.Query(L, R, mid+1, r)
	}
	a := t.left.Query(L, mid, l, mid)
	b := t.right.Query(mid+1, R, mid+1, r)
	return a + b
}
