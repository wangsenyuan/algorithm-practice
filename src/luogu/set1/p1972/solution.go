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

	drive(reader, writer)
}

func drive(reader *bufio.Reader, writer *bufio.Writer) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	var m int
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}

	ans := solve(a, queries)
	for _, x := range ans {
		fmt.Fprintln(writer, x)
	}
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	m := len(queries)
	at := make([][][]int, n)

	for i := range m {
		l, r := queries[i][0]-1, queries[i][1]-1
		at[r] = append(at[r], []int{l, i})
	}

	cnt := make(BIT, n+3)

	pos := make(map[int]int)

	ans := make([]int, m)

	for i := range n {
		if j, ok := pos[a[i]]; ok {
			cnt.Update(j, -1)
		}
		cnt.Update(i, 1)
		pos[a[i]] = i
		for _, q := range at[i] {
			l, id := q[0], q[1]
			ans[id] = cnt.Query(l)
		}
	}
	return ans
}

type BIT []int

func (this BIT) Update(i int, v int) {
	i++
	for i > 0 {
		this[i] += v
		i -= i & -i
	}
}

func (this BIT) Query(i int) int {
	var res int
	i++
	for i < len(this) {
		res += this[i]
		i += i & -i
	}
	return res
}
