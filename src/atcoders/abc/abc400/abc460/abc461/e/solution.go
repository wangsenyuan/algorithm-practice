package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for _, ans := range drive(reader) {
		fmt.Println(ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	queries := make([][]int, q)
	for i := range q {
		var t int
		fmt.Fscan(reader, &t)
		var x int
		fmt.Fscan(reader, &x)
		queries[i] = []int{t, x}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []int {
	s1 := make(BIT, len(queries)+3)
	s2 := make(BIT, len(queries)+3)

	lastRow := make([]int, n)
	lastCol := make([]int, n)

	for i := range n {
		lastRow[i] = -1
		lastCol[i] = -1
	}

	var tot int

	res := make([]int, len(queries))

	for i, cur := range queries {
		if cur[0] == 1 {
			r := cur[1] - 1
			if lastRow[r] >= 0 {
				w := s2.Query(lastRow[r]+1, i)
				tot += w
				s1.Update(lastRow[r], -1)
			} else {
				tot += n
			}
			lastRow[r] = i
			s1.Update(lastRow[r], 1)
		} else {
			c := cur[1] - 1
			w := s1.Query(lastCol[c]+1, i)
			tot -= w
			if lastCol[c] >= 0 {
				s2.Update(lastCol[c], -1)
			}
			lastCol[c] = i
			s2.Update(lastCol[c], 1)
		}

		res[i] = tot
	}

	return res
}

type BIT []int

func (bit BIT) Update(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) Get(p int) int {
	var res int
	p++
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

func (bit BIT) Query(l int, r int) int {
	return bit.Get(r) - bit.Get(l-1)
}
