package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var n, m, k int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m, &k)
	res := solve(n, m, k)

	var buf bytes.Buffer
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d", len(cur)))
		for _, p := range cur {
			buf.WriteString(fmt.Sprintf(" %d %d", p[0], p[1]))
		}
		buf.WriteByte('\n')
	}

	buf.WriteTo(os.Stdout)
}

func solve(n int, m int, k int) [][][]int {
	var res [][][]int

	var cur [][]int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := 0; j < m; j++ {
				cur = append(cur, []int{i + 1, j + 1})
				if len(cur) == 2 && len(res)+1 < k {
					res = append(res, cur)
					cur = make([][]int, 0, 2)
				}
			}
		} else {
			for j := m - 1; j >= 0; j-- {
				cur = append(cur, []int{i + 1, j + 1})
				if len(cur) == 2 && len(res)+1 < k {
					res = append(res, cur)
					cur = make([][]int, 0, 2)
				}
			}
		}
	}

	res = append(res, cur)

	return res
}
