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
		if res {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

func solve(p []int) bool {
	// n := len(p)

	var que []int

	for _, v := range p {
		fa := v
		for len(que) > 0 && que[len(que)-1] < v {
			fa = min(fa, que[len(que)-1])
			que = que[:len(que)-1]
		}
		// 把最小的放进去
		que = append(que, fa)
	}

	return len(que) == 1
}
