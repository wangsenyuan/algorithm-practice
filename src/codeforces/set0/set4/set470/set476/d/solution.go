package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)
	min_num, res := solve(n, k)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", min_num))
	for _, cur := range res {
		for _, num := range cur {
			buf.WriteString(fmt.Sprintf("%d ", num))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
}

const N = 1e7

var fs = []int{1, 2, 3, 5}

func solve(n int, k int) (min_num int, res [][]int) {
	// 先找出4 * n 个质数数出来
	min_num = (6*(n-1) + 5) * k

	for i := range n {
		cur := make([]int, 4)
		for j := range 4 {
			cur[j] = (6*i + fs[j]) * k
		}

		res = append(res, cur)
	}

	return
}
