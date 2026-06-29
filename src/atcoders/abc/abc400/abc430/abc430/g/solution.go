package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans[0], ans[1])
	}
}

func drive(reader *bufio.Reader) [][2]int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	queries := make([][]int, q)
	for i := range q {
		var t int
		fmt.Fscan(reader, &t)
		queries[i] = []int{t}
		if t == 3 {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			queries[i] = append(queries[i], l, r)
		} else {
			var l, r, x int
			fmt.Fscan(reader, &l, &r, &x)
			queries[i] = append(queries[i], l, r, x)
		}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) [][2]int {
	seg := buildSeg(n)
	roots := make([]*interval, 60)
	var res [][2]int
	for _, cur := range queries {
		l, r := cur[1]-1, cur[2]-1
		if cur[0] == 1 {
			x := cur[3] - 1
			addInterval(&roots[x], seg, l, r)
		} else if cur[0] == 2 {
			x := cur[3] - 1
			removeInterval(&roots[x], seg, l, r)
		} else {
			res = append(res, seg.query(l, r))
		}
	}
	return res
}

type segNode struct {
	best int
	cnt  int
	lazy int
}

type segTree struct {
	arr []segNode
	n   int
}

func buildSeg(n int) *segTree {
	tr := &segTree{arr: make([]segNode, 4*n), n: n}
	var build func(int, int, int)
	build = func(i int, l int, r int) {
		tr.arr[i].cnt = r - l + 1
		if l == r {
			return
		}
		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
	}
	build(0, 0, n-1)
	return tr
}

func (tr *segTree) apply(i int, v int) {
	tr.arr[i].best += v
	tr.arr[i].lazy += v
}

func (tr *segTree) pull(i int) {
	a := tr.arr[i*2+1]
	b := tr.arr[i*2+2]
	tr.arr[i].best = max(a.best, b.best)
	tr.arr[i].cnt = 0
	if a.best == tr.arr[i].best {
		tr.arr[i].cnt += a.cnt
	}
	if b.best == tr.arr[i].best {
		tr.arr[i].cnt += b.cnt
	}
}

func (tr *segTree) push(i int) {
	if tr.arr[i].lazy != 0 {
		tr.apply(i*2+1, tr.arr[i].lazy)
		tr.apply(i*2+2, tr.arr[i].lazy)
		tr.arr[i].lazy = 0
	}
}

func (tr *segTree) add(L int, R int, v int) {
	if L > R {
		return
	}
	var dfs func(int, int, int)
	dfs = func(i int, l int, r int) {
		if L <= l && r <= R {
			tr.apply(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			dfs(i*2+1, l, mid)
		}
		if mid < R {
			dfs(i*2+2, mid+1, r)
		}
		tr.pull(i)
	}
	dfs(0, 0, tr.n-1)
}

func mergeAns(a [2]int, b [2]int) [2]int {
	if a[0] > b[0] {
		return a
	}
	if a[0] < b[0] {
		return b
	}
	return [2]int{a[0], a[1] + b[1]}
}

func (tr *segTree) query(L int, R int) [2]int {
	var dfs func(int, int, int) [2]int
	dfs = func(i int, l int, r int) [2]int {
		if L <= l && r <= R {
			return [2]int{tr.arr[i].best, tr.arr[i].cnt}
		}
		tr.push(i)
		mid := (l + r) >> 1
		res := [2]int{-1, 0}
		if L <= mid {
			res = mergeAns(res, dfs(i*2+1, l, mid))
		}
		if mid < R {
			res = mergeAns(res, dfs(i*2+2, mid+1, r))
		}
		return res
	}
	return dfs(0, 0, tr.n-1)
}

type interval struct {
	l, r        int
	priority    uint32
	left, right *interval
}

var seed uint32 = 123456789

func nextRand() uint32 {
	seed ^= seed << 13
	seed ^= seed >> 17
	seed ^= seed << 5
	return seed
}

func newInterval(l int, r int) *interval {
	return &interval{l: l, r: r, priority: nextRand()}
}

func rotate(root *interval, dir int) *interval {
	var x *interval
	if dir == 0 {
		x = root.right
		root.right = x.left
		x.left = root
	} else {
		x = root.left
		root.left = x.right
		x.right = root
	}
	return x
}

func insert(root *interval, cur *interval) *interval {
	if root == nil {
		return cur
	}
	if cur.l < root.l {
		root.left = insert(root.left, cur)
		if root.left.priority < root.priority {
			root = rotate(root, 1)
		}
	} else {
		root.right = insert(root.right, cur)
		if root.right.priority < root.priority {
			root = rotate(root, 0)
		}
	}
	return root
}

func remove(root *interval, key int) *interval {
	if root == nil {
		return nil
	}
	if key < root.l {
		root.left = remove(root.left, key)
	} else if root.l < key {
		root.right = remove(root.right, key)
	} else {
		return mergeInterval(root.left, root.right)
	}
	return root
}

func mergeInterval(a *interval, b *interval) *interval {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.priority < b.priority {
		a.right = mergeInterval(a.right, b)
		return a
	}
	b.left = mergeInterval(a, b.left)
	return b
}

func lowerBound(root *interval, key int) *interval {
	var res *interval
	for root != nil {
		if root.l >= key {
			res = root
			root = root.left
		} else {
			root = root.right
		}
	}
	return res
}

func upperBound(root *interval, key int) *interval {
	var res *interval
	for root != nil {
		if root.l <= key {
			res = root
			root = root.right
		} else {
			root = root.left
		}
	}
	return res
}

func addInterval(root **interval, seg *segTree, L int, R int) {
	nl, nr := L, R
	pos := L
	pre := upperBound(*root, L)
	if pre != nil && pre.r >= L-1 {
		nl = min(nl, pre.l)
		nr = max(nr, pre.r)
		pos = max(pos, pre.r+1)
		*root = remove(*root, pre.l)
	}

	for {
		cur := lowerBound(*root, pos)
		if cur == nil || cur.l > R+1 {
			break
		}
		if pos < cur.l {
			seg.add(pos, min(R, cur.l-1), 1)
		}
		pos = max(pos, cur.r+1)
		nr = max(nr, cur.r)
		key := cur.l
		*root = remove(*root, key)
	}
	if pos <= R {
		seg.add(pos, R, 1)
	}
	*root = insert(*root, newInterval(nl, nr))
}

func removeInterval(root **interval, seg *segTree, L int, R int) {
	cur := upperBound(*root, L)
	if cur == nil || cur.r < L {
		cur = lowerBound(*root, L)
	}

	for cur != nil && cur.l <= R {
		ol, or := cur.l, cur.r
		key := cur.l
		*root = remove(*root, key)
		if max(ol, L) <= min(or, R) {
			seg.add(max(ol, L), min(or, R), -1)
		}
		if ol < L {
			*root = insert(*root, newInterval(ol, L-1))
		}
		if R < or {
			*root = insert(*root, newInterval(R+1, or))
		}
		cur = lowerBound(*root, max(L, key+1))
	}
}
