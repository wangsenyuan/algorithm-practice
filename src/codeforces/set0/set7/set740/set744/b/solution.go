package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := make([]int, n)
	ask := func(w []int) []int {
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "%d\n", len(w))
		for _, v := range w {
			fmt.Fprintf(&buf, "%d ", v)
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		for i := range n {
			fmt.Fscan(reader, &res[i])
		}
		return res
	}
	ans := solve(n, ask)
	var buf bytes.Buffer
	buf.WriteString("-1\n")
	for _, v := range ans {
		fmt.Fprintf(&buf, "%d ", v)
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

const inf = 1 << 60

func solve(n int, ask func(w []int) []int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = inf
	}

	for d := range 10 {
		for v := range 2 {
			var w []int
			for i := 1; i <= n; i++ {
				if (i>>d)&1 == v {
					w = append(w, i)
				}
			}
			if len(w) > 0 && len(w) < n {
				res := ask(w)
				for i := 1; i <= n; i++ {
					if (i>>d)&1 == v^1 {
						ans[i] = min(ans[i], res[i-1])
					}
				}
			}
		}
	}

	return ans[1:]
}
