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
		var n, l, r int
		fmt.Fscan(reader, &n, &l, &r)
		res := solve(n, l, r)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func solve(n int, l int, r int) []int {
	res := make([]int, n)

	get := func(i int) int {
		if i == r {
			return l - 1
		}
		return i
	}

	for i := range n {
		res[i] = get(i) ^ get(i+1)
	}
	return res
}
