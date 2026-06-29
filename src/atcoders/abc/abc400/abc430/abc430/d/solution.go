package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var n int
	fmt.Fscan(reader, &n)
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	return solve(x)
}

func solve(x []int) []int64 {
	t := new(rbtree)

	t.insert(0)

	var sum int64

	var res []int64

	for _, v := range x {
		// 先找到v左边的
		// l 是最大的x < v
		l := t.upperBound(v - 1)

		// r > v, 最小的 x > v
		r := t.lowerBound(v + 1)

		// 考虑l的情况
		if l != nil {
			l1 := t.upperBound(l.data - 1)
			if l1 != nil && r != nil {
				sum -= int64(min(l.data-l1.data, r.data-l.data))
			} else if l1 != nil {
				sum -= int64(l.data - l1.data)
			} else if r != nil {
				sum -= int64(r.data - l.data)
			}
			if l1 != nil {
				sum += int64(min(l.data-l1.data, v-l.data))
			} else {
				sum += int64(v - l.data)
			}
		}

		if r != nil {
			r1 := t.lowerBound(r.data + 1)
			if l != nil && r1 != nil {
				sum -= int64(min(r.data-l.data, r1.data-r.data))
			} else if l != nil {
				sum -= int64(r.data - l.data)
			} else if r1 != nil {
				sum -= int64(r1.data - r.data)
			}

			if r1 != nil {
				sum += int64(min(r.data-v, r1.data-r.data))
			} else {
				sum += int64(r.data - v)
			}
		}

		if l != nil && r != nil {
			sum += int64(min(v-l.data, r.data-v))
		} else if l != nil {
			sum += int64(v - l.data)
		} else if r != nil {
			sum += int64(r.data - v)
		}

		res = append(res, sum)
		t.insert(v)
	}

	return res
}

type node struct {
	fa   *node
	ch   [2]*node
	data int
	red  bool // 0 or 1
	sz   int
}

func newNode(data int) *node {
	n := new(node)
	n.sz = 1
	n.data = data
	n.red = true
	return n
}

func (node *node) size() int {
	if node == nil {
		return 0
	}
	return node.sz
}

type rbtree struct {
	root *node
}

func (t *rbtree) rotate(p *node, dir int) *node {
	g := p.fa
	s := p.ch[1^dir]
	s.sz = p.sz

	p.sz = p.ch[dir].size() + s.ch[dir].size() + 1
	c := s.ch[dir]
	if c != nil {
		c.fa = p
	}
	p.ch[dir^1] = c
	s.ch[dir] = p

	p.fa = s
	s.fa = g

	if g != nil {
		if g.ch[1] == p {
			g.ch[1] = s
		} else {
			g.ch[0] = s
		}
	} else {
		t.root = s
	}

	return s
}

func (t *rbtree) insert(data int) *node {
	n := newNode(data)
	now := t.root
	var parent *node
	var dir int
	for now != nil {
		parent = now
		now.sz++
		if data < now.data {
			dir = 0
		} else {
			dir = 1
		}
		now = now.ch[dir]
	}
	n.fa = parent
	if parent == nil {
		t.root = n
		t.root.red = false
		return n
	}

	parent.ch[dir] = n

	t.fixInsert(n)
	return n
}

func (t *rbtree) fixInsert(n *node) {
	// n.fa.red => n.fa.fa != nil
	for n != t.root && n.fa.red {
		var dir int
		if n.fa == n.fa.fa.ch[1] {
			dir = 1
		}
		uncle := n.fa.fa.ch[1^dir]
		if uncle != nil && uncle.red {
			n.fa.red = false
			uncle.red = false
			n.fa.fa.red = true
			n = n.fa.fa
		} else {
			if n == n.fa.ch[1^dir] {
				n = n.fa
				t.rotate(n, dir)
			}
			n.fa.red = false
			n.fa.fa.red = true
			t.rotate(n.fa.fa, 1^dir)
		}
	}

	t.root.red = false
}

func (t *rbtree) upperBound(data int) *node {
	now := t.root
	var res *node
	for now != nil {
		if now.data <= data {
			res = now
			now = now.ch[1]
		} else {
			// now.data >= data
			now = now.ch[0]
		}
	}
	return res
}

func (t *rbtree) lowerBound(data int) *node {
	now := t.root
	var res *node
	for now != nil {
		if now.data >= data {
			res = now
			now = now.ch[0]
		} else {
			now = now.ch[1]
		}
	}
	return res
}
