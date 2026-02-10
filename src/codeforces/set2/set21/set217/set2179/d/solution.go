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
		for i, v := range res {
			if i < len(res)-1 {
				fmt.Fprintf(writer, "%d ", v)
			} else {
				fmt.Fprintf(writer, "%d\n", v)
			}
		}
	}
}

func solve(n int) []int {
	var res []int

	res = append(res, 1<<n-1)

	for d := n - 1; d >= 0; d-- {
		v := 1<<d - 1
		for j := range 1 << (n - 1 - d) {
			res = append(res, j<<(d+1)|v)
		}
	}

	return res
}
