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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	freq := make(map[int]int)

	dp := make([]int, n+1)

	bit := make(BIT, n+3)

	for i, v := range a {
		freq[v]++
		dp[i] = freq[v]
		bit.update(dp[i], 1)
	}

	clear(freq)

	var res int

	for i := n - 1; i >= 0; i-- {
		// 先取消i的贡献
		bit.update(dp[i], -1)
		v := a[i]
		freq[v]++
		w := freq[v]
		res += bit.queryRange(w+1, n)
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

func (bit BIT) query(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) queryRange(l, r int) int {
	return bit.query(r) - bit.query(l-1)
}
