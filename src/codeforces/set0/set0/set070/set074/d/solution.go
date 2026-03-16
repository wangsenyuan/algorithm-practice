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
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	queries := make([][]int, q)
	for i := range q {
		var k int
		fmt.Fscan(reader, &k)
		if k == 0 {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			queries[i] = []int{0, l, r}
		} else {
			queries[i] = []int{k}
		}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []int {
	tr := NewNode(0, n-1)

	var ans []int

	pos := make(map[int]int)

	for _, cur := range queries {
		if cur[0] == 0 {
			l, r := cur[1]-1, cur[2]-1
			ans = append(ans, tr.Count(l, r))
		} else {
			x := cur[0]
			if j, ok := pos[x]; ok {
				tr.Update(j, 0)
				delete(pos, x)
			} else {
				l, r := tr.GetBestPosition(0, n-1)
				mid := (l + r + 1) / 2
				tr.Update(mid, 1)
				pos[x] = mid
			}
		}
	}

	return ans
}

const inf = 1 << 60

type node struct {
	l, r                int
	leftMost, rightMost int
	maxDist             int
	cnt                 int
	left, right         *node
}

func NewNode(l int, r int) *node {
	return &node{l: l, r: r, leftMost: inf, rightMost: -1, maxDist: r - l + 1, cnt: 0}
}

func (n *node) isEmpty() bool {
	return n == nil || n.cnt == 0
}

func (n *node) merge(lc *node, rc *node) {
	if lc.isEmpty() && rc.isEmpty() {
		n.maxDist = n.r - n.l + 1
		n.cnt = 0
		return
	}
	if lc.isEmpty() {
		n.maxDist = rc.leftMost - n.l
		n.cnt = rc.cnt
		n.leftMost = rc.leftMost
		n.rightMost = rc.rightMost
		return
	}
	if rc.isEmpty() {
		n.maxDist = n.r - lc.rightMost
		n.cnt = lc.cnt
		n.leftMost = lc.leftMost
		n.rightMost = lc.rightMost
		return
	}
	n.maxDist = max(lc.maxDist, rc.maxDist, rc.leftMost-lc.rightMost-1)
	n.cnt = lc.cnt + rc.cnt
	n.leftMost = lc.leftMost
	n.rightMost = rc.rightMost
}

func (n *node) Update(p int, v int) {
	if n.l == n.r {
		if v == 1 {
			n.leftMost = n.l
			n.rightMost = n.l
			n.cnt = 1
			n.maxDist = 0
		} else {
			n.leftMost = inf
			n.rightMost = -1
			n.cnt = 0
			n.maxDist = 1
		}
		return
	}
	mid := (n.l + n.r) >> 1
	if p <= mid {
		if n.left == nil {
			n.left = NewNode(n.l, mid)
		}
		n.left.Update(p, v)
	} else {
		if n.right == nil {
			n.right = NewNode(mid+1, n.r)
		}
		n.right.Update(p, v)
	}
	n.merge(n.left, n.right)
}

func (n *node) GetBestPosition(l int, r int) (int, int) {
	if n.isEmpty() {
		return l, r
	}
	if n.left.isEmpty() {
		return l, n.right.leftMost - 1
	}
	if n.right.isEmpty() {
		return n.left.rightMost + 1, r
	}
	mid := (l + r) >> 1
	if n.maxDist == n.right.maxDist {
		return n.right.GetBestPosition(mid+1, r)
	}
	if n.maxDist == n.right.leftMost-n.left.rightMost-1 {
		return n.left.rightMost + 1, n.right.leftMost - 1
	}
	return n.left.GetBestPosition(l, mid)
}

func (n *node) Count(L int, R int) int {
	if n.isEmpty() || R < n.l || n.r < L {
		return 0
	}
	if L <= n.l && n.r <= R {
		return n.cnt
	}
	return n.left.Count(L, R) + n.right.Count(L, R)
}
