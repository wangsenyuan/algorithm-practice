package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)

	for _, x := range res {
		fmt.Println(x)
	}
}

type query struct {
	gid int
	s   string
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	queries := make([]query, n)
	for i := range n {
		var qry query
		fmt.Fscan(reader, &qry.gid, &qry.s)
		queries[i] = qry
	}
	return solve(queries)
}

func solve(queries []query) []int {
	// 加入一个 x的时候，如果现有的答案是tot, 且x是w个y中字符串的前缀, tot -= w
	// 加入一个 y, 如果y是包含任何一个x作为前缀, tot不变，否则tot++
	tx := NewTrie()
	ty := NewTrie()

	var tot int

	add1 := func(s string) {
		var cur int
		for _, x := range s {
			d := int(x - 'a')
			if ty.next[cur][d] == 0 {
				cur = -1
				break
			}
			cur = ty.next[cur][d]
		}
		if cur > 0 {
			w := ty.cnt[cur]
			tot -= w
			var cur1 int
			for _, x := range s {
				d := int(x - 'a')
				if ty.next[cur1][d] == cur {
					// 后面的部分全部删掉
					ty.next[cur1][d] = 0
					break
				}
				// 它的父节点全部 -w
				cur1 = ty.next[cur1][d]
				ty.cnt[cur1] -= w
			}
		}
		tx.add(s)
	}

	add2 := func(s string) {
		var cur int
		for _, x := range s {
			d := int(x - 'a')
			if tx.next[cur][d] == 0 {
				// not a prefix
				cur = -1
				break
			}
			cur = tx.next[cur][d]
			if tx.leaf[cur] {
				// found a prefix
				return
			}
		}

		tot++
		ty.add(s)
	}

	res := make([]int, len(queries))

	for i, cur := range queries {
		if cur.gid == 1 {
			add1(cur.s)
		} else {
			add2(cur.s)
		}
		res[i] = tot
	}

	return res
}

type trie struct {
	next [][26]int
	cnt  []int
	leaf []bool
}

func NewTrie() *trie {
	return &trie{
		next: make([][26]int, 1),
		cnt:  make([]int, 1),
		leaf: make([]bool, 1),
	}
}

func (tr *trie) expand() int {
	tr.next = append(tr.next, [26]int{})
	tr.cnt = append(tr.cnt, 0)
	tr.leaf = append(tr.leaf, false)
	return len(tr.next) - 1
}

func (tr *trie) add(s string) {
	var cur int
	for _, x := range s {
		d := int(x - 'a')
		if tr.next[cur][d] == 0 {
			tr.next[cur][d] = tr.expand()
		}
		cur = tr.next[cur][d]
		tr.cnt[cur]++
	}
	tr.leaf[cur] = true
}
