package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
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
		var p, k int
		fmt.Fscan(reader, &p, &k)
		queries[i] = []int{p, k}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	sq := int(math.Sqrt(float64(n)))

	small := make([][]int, sq+1)

	ans := make([]int, len(queries))

	bruteForce := func(p int, k int) int {
		p--
		var res int
		for p < n {
			res++
			p += a[p] + k
		}
		return res
	}

	for i, cur := range queries {
		k := cur[1]
		if k <= sq {
			small[k] = append(small[k], i)
		} else {
			ans[i] = bruteForce(cur[0], cur[1])
		}
	}

	dp := make([]int, n)

	for k := 1; k <= sq; k++ {
		for i := n - 1; i >= 0; i-- {
			if i+a[i]+k >= n {
				dp[i] = 1
			} else {
				dp[i] = dp[i+a[i]+k] + 1
			}
		}
		for _, i := range small[k] {
			p := queries[i][0] - 1
			ans[i] = dp[p]
		}
	}

	return ans
}
