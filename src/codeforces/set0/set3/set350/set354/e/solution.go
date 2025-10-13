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
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n)
		if res == nil {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintln(writer, res[0], res[1], res[2], res[3], res[4], res[5])
		}
	}
}

type pair struct {
	first  int
	second int
}

var fp [10][]int
var dp [100]pair

func init() {
	for j := range 7 {
		for i := 0; i+j <= 6; i++ {
			r := (j*4 + i*7)
			fp[r%10] = append(fp[r%10], r)
			dp[r] = pair{j, i}
		}
	}
}

func solve(n int) []int {
	s := fmt.Sprintf("%d", n)
	m := len(s)
	ans := make([][]int, 6)
	for i := range 6 {
		ans[i] = make([]int, m)
		for j := range m {
			ans[i][j] = 0
		}
	}

	vis := make([][]bool, m)
	for i := range m {
		vis[i] = make([]bool, 7)
	}
	var dfs func(pos int, c int) bool

	dfs = func(pos int, c int) bool {
		if pos < 0 {
			return c == 0
		}

		if vis[pos][c] {
			return false
		}
		vis[pos][c] = true
		cur := int(s[pos] - '0')
		cur -= c
		r := cur % 10
		if r < 0 {
			r += 10
		}
		for _, sum := range fp[r] {
			rest := cur - sum
			if rest > 0 || rest%10 != 0 {
				continue
			}
			rest /= 10
			if !dfs(pos-1, -rest) {
				continue
			}
			c4, c7 := dp[sum].first, dp[sum].second
			for i := range c4 {
				ans[i][pos] = 4
			}
			for i := c4; i < c4+c7; i++ {
				ans[i][pos] = 7
			}
			return true
		}

		return false
	}

	if !dfs(m-1, 0) {
		return nil
	}

	res := make([]int, 6)
	for i := range 6 {
		for j := range m {
			res[i] = res[i]*10 + ans[i][j]
		}
	}
	return res
}
