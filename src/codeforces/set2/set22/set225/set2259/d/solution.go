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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) string {
	var zeros []int
	for i, v := range a {
		if v == 0 {
			zeros = append(zeros, i)
		}
	}
	if len(zeros) == 1 {
		return "NO"
	}
	buf := make([]byte, len(a))
	if len(zeros) > 0 {
		buf[zeros[0]] = 'A'
		for i := 1; i < len(zeros); i++ {
			buf[zeros[i]] = 'B'
		}
	}
	for i, v := range a {
		if v != 0 {
			buf[i] = 'C'
		}
	}
	return fmt.Sprintf("YES\n%s", string(buf))
}
