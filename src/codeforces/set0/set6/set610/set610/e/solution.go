package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := solve(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

type state struct {
	first int
	last  int
	cnt   [][]int
}

func newMatrix(k int) [][]int {
	m := make([][]int, k)
	for i := range k {
		m[i] = make([]int, k)
	}
	return m
}

func mergeState(a state, b state, k int) state {
	c := newMatrix(k)

	for i := range k {
		for j := range k {
			c[i][j] = a.cnt[i][j] + b.cnt[i][j]
		}
	}
	c[a.last][b.first]++
	return state{
		first: a.first,
		last:  b.last,
		cnt:   c,
	}
}

func solve(reader *bufio.Reader) []int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	var s string
	fmt.Fscan(reader, &s)

	tr := make([]state, 4*n)
	lazy := make([]int, 4*n)

	merge := func(a state, b state, c *state) {
		for u := range k {
			for v := range k {
				c.cnt[u][v] = a.cnt[u][v] + b.cnt[u][v]
			}
		}
		c.first = a.first
		c.last = b.last
		c.cnt[a.last][b.first]++
	}

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		lazy[i] = -1
		if l == r {
			tr[i] = state{
				first: int(s[l] - 'a'),
				last:  int(s[l] - 'a'),
				cnt:   newMatrix(k),
			}
			return
		}
		mid := (l + r) / 2
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		tr[i].cnt = newMatrix(k)
		merge(tr[i*2+1], tr[i*2+2], &tr[i])
	}

	build(0, 0, n-1)

	apply := func(i int, l int, r int, c int) {
		for u := range k {
			for v := range k {
				tr[i].cnt[u][v] = 0
			}
		}
		tr[i].first = c
		tr[i].last = c
		tr[i].cnt[c][c] = r - l
		lazy[i] = c
	}

	push := func(i int, l int, r int) {
		if lazy[i] != -1 {
			mid := (l + r) / 2
			apply(i*2+1, l, mid, lazy[i])
			apply(i*2+2, mid+1, r, lazy[i])
			lazy[i] = -1
		}
	}

	var update func(i int, l int, r int, L int, R int, c int)
	update = func(i int, l int, r int, L int, R int, c int) {
		if l == L && r == R {
			apply(i, l, r, c)
			return
		}
		push(i, l, r)
		mid := (l + r) / 2
		if L <= mid {
			update(i*2+1, l, mid, L, min(mid, R), c)
		}
		if mid < R {
			update(i*2+2, mid+1, r, max(mid+1, L), R, c)
		}
		merge(tr[i*2+1], tr[i*2+2], &tr[i])
	}

	pos := make([]int, k)
	find := func(p string) int {
		for i, x := range p {
			pos[int(x-'a')] = i
		}
		var sum int
		for i := range k {
			for j := range k {
				c := tr[0].cnt[i][j]
				if pos[i] >= pos[j] {
					sum += c
				}
			}
		}
		return sum + 1
	}

	var ans []int

	for range m {
		var tp int
		fmt.Fscan(reader, &tp)
		if tp == 1 {
			var l, r int
			var c string
			fmt.Fscan(reader, &l, &r, &c)
			update(0, 0, n-1, l-1, r-1, int(c[0]-'a'))
		} else {
			var p string
			fmt.Fscan(reader, &p)
			ans = append(ans, find(p))
		}
	}

	return ans
}

func drive(reader *bufio.Reader) []int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	var s string
	fmt.Fscan(reader, &s)

	st := NewSegTree([]byte(s), k)
	var res []int
	for ; m > 0; m-- {
		var tp int
		fmt.Fscan(reader, &tp)
		if tp == 1 {
			var l, r int
			var c string
			fmt.Fscan(reader, &l, &r, &c)
			st.Update(l-1, r-1, int(c[0]-'a'))
		} else {
			var p string
			fmt.Fscan(reader, &p)
			res = append(res, st.Query([]byte(p)))
		}
	}
	return res
}

const maxAlphabet = 10

type Node struct {
	l     int
	r     int
	c     int
	pri   uint32
	left  *Node
	right *Node
	first int
	last  int
	cnt   [maxAlphabet * maxAlphabet]int32
}

type SegTree struct {
	k    int
	root *Node
	seed uint32
}

func NewSegTree(s []byte, k int) *SegTree {
	st := &SegTree{
		k:    k,
		seed: 1,
	}
	for l := 0; l < len(s); {
		r := l
		for r+1 < len(s) && s[r+1] == s[l] {
			r++
		}
		st.root = merge(st.root, st.newNode(l, r, int(s[l]-'a')))
		l = r + 1
	}
	return st
}

func (st *SegTree) newNode(l int, r int, c int) *Node {
	node := &Node{
		l:     l,
		r:     r,
		c:     c,
		pri:   st.nextRand(),
		first: c,
		last:  c,
	}
	node.cnt[c*maxAlphabet+c] = int32(r - l)
	return node
}

func (st *SegTree) nextRand() uint32 {
	st.seed = st.seed*1664525 + 1013904223
	return st.seed
}

func (st *SegTree) Update(l int, r int, c int) {
	var mid, right *Node
	st.root, mid = st.split(st.root, l)
	mid, right = st.split(mid, r+1)
	st.root = merge(merge(st.root, st.newNode(l, r, c)), right)
}

func (st *SegTree) Query(p []byte) int {
	pos := make([]int, st.k)
	for i, x := range p {
		pos[x-'a'] = i
	}
	res := 1
	for i := 0; i < st.k; i++ {
		for j := 0; j < st.k; j++ {
			if pos[i] >= pos[j] {
				res += int(st.root.cnt[i*maxAlphabet+j])
			}
		}
	}
	return res
}

func (st *SegTree) split(root *Node, pos int) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}
	if pos <= root.l {
		a, b := st.split(root.left, pos)
		root.left = b
		pull(root)
		return a, root
	}
	if pos > root.r {
		a, b := st.split(root.right, pos)
		root.right = a
		pull(root)
		return root, b
	}

	left := st.newNode(root.l, pos-1, root.c)
	right := st.newNode(pos, root.r, root.c)
	return merge(root.left, left), merge(right, root.right)
}

func merge(a *Node, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.pri < b.pri {
		a.right = merge(a.right, b)
		pull(a)
		return a
	}
	b.left = merge(a, b.left)
	pull(b)
	return b
}

func pull(node *Node) {
	for i := range node.cnt {
		node.cnt[i] = 0
	}
	has := false
	add := func(first int, last int, cnt [maxAlphabet * maxAlphabet]int32) {
		if !has {
			node.first = first
			has = true
		} else {
			node.cnt[node.last*maxAlphabet+first]++
		}
		for i, v := range cnt {
			node.cnt[i] += v
		}
		node.last = last
	}
	if node.left != nil {
		add(node.left.first, node.left.last, node.left.cnt)
	}
	if !has {
		node.first = node.c
		has = true
	} else {
		node.cnt[node.last*maxAlphabet+node.c]++
	}
	node.cnt[node.c*maxAlphabet+node.c] += int32(node.r - node.l)
	node.last = node.c
	if node.right != nil {
		add(node.right.first, node.right.last, node.right.cnt)
	}
}
