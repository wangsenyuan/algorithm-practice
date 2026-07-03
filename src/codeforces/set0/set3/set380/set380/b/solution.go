package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	ops := make([][]int, m)
	for i := range m {
		var typ int
		fmt.Fscan(reader, &typ)
		if typ == 1 {
			var t, l, r, x int
			fmt.Fscan(reader, &t, &l, &r, &x)
			ops[i] = []int{typ, t, l, r, x}
		} else {
			var t, v int
			fmt.Fscan(reader, &t, &v)
			ops[i] = []int{typ, t, v}
		}
	}
	return solve(n, ops)
}

func solve(n int, ops [][]int) []int {
	type Update struct {
		time    int
		t, l, r int
		x       int
	}
	type Query struct {
		time int
		t, v int
	}

	parent := func(pos int) int {
		res := pos - bits.Len(uint(pos))
		if res < 1 {
			res = 1
		}
		for res+bits.Len(uint(res)) < pos {
			res++
		}
		for res > 1 && res-1+bits.Len(uint(res-1)) >= pos {
			res--
		}
		return res
	}

	id := make(map[int]int)
	var updates []Update
	var queries []Query

	for i, op := range ops {
		if op[0] == 1 {
			x := op[4]
			if _, ok := id[x]; !ok {
				id[x] = len(id)
			}
			updates = append(updates, Update{i, op[1], op[2], op[3], id[x]})
		} else {
			queries = append(queries, Query{i, op[1], op[2]})
		}
	}

	ans := make([]int, len(queries))
	seen := make([][]uint64, len(queries))
	for i := range seen {
		seen[i] = make([]uint64, (len(id)+63)/64)
	}

	L := make([]int, n+1)
	R := make([]int, n+1)
	for _, cur := range updates {
		L[cur.t] = cur.l
		R[cur.t] = cur.r
		for level := cur.t; level > 1; level-- {
			L[level-1] = parent(L[level])
			R[level-1] = parent(R[level])
		}

		for i, q := range queries {
			if q.time <= cur.time || q.t > cur.t {
				continue
			}
			if L[q.t] <= q.v && q.v <= R[q.t] {
				k := cur.x / 64
				b := uint(cur.x % 64)
				if (seen[i][k]>>b)&1 == 0 {
					seen[i][k] |= 1 << b
					ans[i]++
				}
			}
		}
	}

	return ans
}
