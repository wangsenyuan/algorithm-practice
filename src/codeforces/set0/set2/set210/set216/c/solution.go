package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, m, k int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m, &k)
	res := solve(n, m, k)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, v := range res {
		buf.WriteString(fmt.Sprintf("%d ", v))
	}
	buf.WriteByte('\n')
	buf.WriteTo(os.Stdout)
}

func solve(n int, m int, k int) []int {
	// n > 1
	if k == 1 {
		return solve1(n, m)
	}
	var res []int
	for range k {
		res = append(res, 1)
	}
	// 前k个人，必须在第一天开始工作
	day := n + 1
	for day <= n+m {
		// 后面每k-1个人，在day天开始
		for range k - 1 {
			res = append(res, day)
		}
		day += n
	}

	// 然后在第 n - 1, 2 * n - 1, ... 天额外雇佣一个人
	for day := n; day <= n+m; day += n {
		res = append(res, day)
	}
	sort.Ints(res)
	return res
}

func solve1(n int, m int) []int {
	var res []int
	for i := 1; ; i += n - 1 {
		res = append(res, i)
		if i+n-1 > m+n {
			break
		}
	}
	return res
}
