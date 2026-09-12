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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return formatInts(solve(m, a))
}

func formatInts(a []int64) string {
	s := fmt.Sprintf("%v", a)
	return s[1 : len(s)-1]
}

func solve(m int, a []int) []int64 {
	// TODO
	_ = a
	return make([]int64, m)
}
