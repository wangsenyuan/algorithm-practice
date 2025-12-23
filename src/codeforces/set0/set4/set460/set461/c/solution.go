package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
		if queries[i][0] == 2 {
			fmt.Fscan(reader, &queries[i][2])
		}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []int {

	cnt := make(BIT, n+10)

	for i := range n {
		cnt.update(i, 1)
	}

	var revered bool
	l, r := 0, n-1

	a := make([]int, n)
	for i := range n {
		a[i] = 1
	}

	fold := func(d int) {
		width := r - l + 1
		if !revered {
			if d <= width/2 {
				for i := l; i < l+d; i++ {
					j := 2*(l+d) - 1 - i
					a[j] += a[i]
					cnt.update(j, a[i])
				}
				l += d
			} else {
				d = width - d
				for i := r; i > r-d; i-- {
					j := 2*(r-d+1) - 1 - i
					a[j] += a[i]
					cnt.update(j, a[i])
				}
				r -= d
				revered = true
			}
		} else {
			if d <= width/2 {
				for i := r; i > r-d; i-- {
					j := 2*(r-d+1) - 1 - i
					a[j] += a[i]
					cnt.update(j, a[i])
				}
				r -= d
			} else {
				d = width - d
				for i := l; i < l+d; i++ {
					j := 2*(l+d) - 1 - i
					a[j] += a[i]
					cnt.update(j, a[i])
				}
				l += d
				revered = false
			}
		}
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			fold(cur[1])
		} else {
			if !revered {
				l0 := cur[1] + l
				r0 := cur[2] + l - 1
				ans = append(ans, cnt.query(l0, r0))
			} else {
				l0 := r - cur[2] + 1
				r0 := r - cur[1]
				ans = append(ans, cnt.query(l0, r0))
			}
		}
	}

	return ans
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) query(l int, r int) int {
	return bit.get(r) - bit.get(l-1)
}
