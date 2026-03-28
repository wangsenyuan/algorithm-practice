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
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) int {
	n := len(a)
	var res int

	var stack []pair

	for i, v := range a {
		for len(stack) > 0 && stack[len(stack)-1].first >= v {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && stack[len(stack)-1].first+1 == v {
			res += (n - i) * (i - stack[len(stack)-1].second)
		} else {
			res += (n - i) * (i + 1)
			stack = stack[:0]
		}
		stack = append(stack, pair{v, i})
	}
	return res
}
