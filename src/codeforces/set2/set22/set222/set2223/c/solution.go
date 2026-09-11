package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	f := make([]int, n+1)
	for u := 2; u <= n; u++ {
		fmt.Fscan(reader, &f[u])
	}
	l := make([]int, n+1)
	for u := 2; u <= n; u++ {
		fmt.Fscan(reader, &l[u])
	}
	ms := make([]int64, q)
	for i := range q {
		fmt.Fscan(reader, &ms[i])
	}
	return formatInts(solve(f, l, ms))
}

func formatInts(a []int) string {
	var b strings.Builder
	for i, x := range a {
		if i > 0 {
			b.WriteByte(' ')
		}
		fmt.Fprintf(&b, "%d", x)
	}
	return b.String()
}

func solve(f, l []int, ms []int64) []int {
	// TODO
	_ = f
	_ = l
	return make([]int, len(ms))
}
