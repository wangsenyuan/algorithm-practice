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
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n)
		if res == nil {
			fmt.Fprintln(writer, -1)
			continue
		}
		for i, x := range res {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, x)
		}
		fmt.Fprintln(writer)
	}
}

func solve(n int) []int {
	if n&(n-1) == 0 {
		return nil
	}

	p := make([]int, n)
	if n&1 == 1 {
		p[0] = n - 1
		for i := 2; i <= n-2; i++ {
			p[i-1] = i ^ 1
		}
		p[n-2] = 1
		p[n-1] = n
		return p
	}

	d := n ^ (n - 1)
	p[0] = d ^ 1
	for i := 2; i <= n-2; i++ {
		p[i-1] = i ^ 1
	}
	p[d-1] = n
	p[n-2] = 1
	p[n-1] = n - 2
	return p
}
