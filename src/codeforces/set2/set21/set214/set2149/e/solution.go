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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, k, L, R int
	fmt.Fscan(reader, &n, &k, &L, &R)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, L, R, a)
}

func solve(k int, L int, R int, a []int) int {
	n := len(a)
	bit := make(BIT, n+2)

	pos := make(map[int]int)

	var lo, hi int

	var res int

	for i, v := range a {
		if j, ok := pos[v]; ok {
			bit.update(j, -1)
		}
		pos[v] = i
		bit.update(i, 1)

		for bit.rangeSum(lo, i) > k {
			lo++
		}
		for bit.rangeSum(hi, i) >= k {
			hi++
		}
		// i - j + 1 >= L => j <= i - L + 1
		// i - j + 1 <= R => j >= i - R + 1
		if bit.rangeSum(lo, i) == k {
			cnt := min(i-L+2, hi) - max(i-R+1, lo)
			res += max(cnt, 0)
		}
	}
	return res
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
	var res int
	i++
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) rangeSum(l int, r int) int {
	return bit.get(r) - bit.get(l-1)
}
