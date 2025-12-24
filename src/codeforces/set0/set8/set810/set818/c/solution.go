package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var d, n, m int
	fmt.Fscan(reader, &d, &n, &m)
	arr := make([][]int, d)
	for i := range d {
		arr[i] = make([]int, 4)
		for j := range 4 {
			fmt.Fscan(reader, &arr[i][j])
		}
	}
	cnt := make([]int, 4)
	for i := range 4 {
		fmt.Fscan(reader, &cnt[i])
	}
	return solve(n, m, arr, cnt)
}

func solve(n int, m int, sofas [][]int, cnt []int) int {
	k := len(sofas)

	l := make(BIT, n+3)
	r := make(BIT, n+3)
	t := make(BIT, m+3)
	b := make(BIT, m+3)

	for i := range k {
		x1, y1, x2, y2 := sofas[i][0], sofas[i][1], sofas[i][2], sofas[i][3]
		x1, x2 = min(x1, x2), max(x1, x2)
		y1, y2 = min(y1, y2), max(y1, y2)

		l.update(x1, 1)
		r.update(x2, 1)
		t.update(y1, 1)
		b.update(y2, 1)
	}

	for i := range k {
		x1, y1, x2, y2 := sofas[i][0], sofas[i][1], sofas[i][2], sofas[i][3]
		x1, x2 = min(x1, x2), max(x1, x2)
		y1, y2 = min(y1, y2), max(y1, y2)

		l.update(x1, -1)
		r.update(x2, -1)
		t.update(y1, -1)
		b.update(y2, -1)

		c0 := l.queryRange(0, x2-1)
		c1 := r.queryRange(x1+1, n)
		c2 := t.queryRange(0, y2-1)
		c3 := b.queryRange(y1+1, m)

		if c0 == cnt[0] && c1 == cnt[1] && c2 == cnt[2] && c3 == cnt[3] {
			return i + 1
		}

		l.update(x1, 1)
		r.update(x2, 1)
		t.update(y1, 1)
		b.update(y2, 1)
	}

	return -1
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) query(i int) int {
	var res int
	i++
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) queryRange(l int, r int) int {
	if r < l {
		return 0
	}
	return bit.query(r) - bit.query(l-1)
}
