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

	res := drive(reader)
	fmt.Fprintln(writer, res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	h := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &h[i])
	}
	return solve(h)
}

func solve(h []int) int {
	n := len(h)
	// 如果h[i] <= min(suf), 那么就可以在i处切割
	suf := make([]int, n+1)
	suf[n] = 1 << 60
	for i := n - 1; i >= 0; i-- {
		suf[i] = min(suf[i+1], h[i])
	}
	var pref int
	var res int
	for i := range n {
		pref = max(pref, h[i])
		if pref <= suf[i+1] {
			res++
			pref = 0
		}
	}
	return res
}
