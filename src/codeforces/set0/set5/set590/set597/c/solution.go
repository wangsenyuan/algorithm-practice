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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) int {
	// dp[i][j] 不大行，因为必须找到比自己小的数
	n := len(a)
	// pos[0] = 0
	k++
	dp := make([]BIT, k+1)
	for i := range k + 1 {
		dp[i] = make(BIT, n+2)
	}
	dp[0].update(0, 1)

	for i := 1; i <= n; i++ {
		for j := k; j > 0; j-- {
			v := a[i-1]
			dp[j].update(v, dp[j-1].query(v-1))
		}
	}
	return dp[k].query(n)
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
