package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var n, y int
	fmt.Fscan(reader, &n, &y)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	return solve(c, y)
}

func solve(c []int, y int) int {
	mx := slices.Max(c)
	freq := make([]int, mx+10)
	for _, v := range c {
		freq[v]++
	}

	// c[i] = 1 的比较特殊，无论多少它都是满足条件的
	n := len(c) - freq[1]

	special := freq[1]

	if n == 0 {
		return special
	}

	freq[1] = 0
	for i := 1; i <= mx; i++ {
		freq[i] += freq[i-1]
	}

	dp := make([]int, mx+1)

	fp := make([]int, mx+1)

	for i := 1; i <= mx; i++ {
		// 有w个i
		w := freq[i] - freq[i-1]
		ed := (mx + i - 1) / i * i
		for j := 2 * i; j <= ed; j += i {
			x := j / i
			if (i-1)*x > mx {
				break
			}
			v := freq[min(mx, j)] - freq[(i-1)*x]
			// i对x的贡献 = min(w, v)
			z := min(w, v)
			// 无论怎么样，需要有v个价格为i的标签，其中z个可以使用现有的w个代替
			// 其他的 v- z 个需要额外供应
			dp[x] += v*i - (v-z)*y
			fp[x] += v
		}
	}
	ans := -(1 << 60)
	for x := 2; x <= mx; x++ {
		if fp[x] == n {
			ans = max(ans, dp[x])
		}
	}
	// 不包括 x = 0/1
	return ans + special
}
