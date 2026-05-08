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
		var n, m, k int
		fmt.Fscan(reader, &n, &m, &k)
		res := solve(n, m, k)
		for i := range n {
			for j := range m {
				fmt.Fprint(writer, res[i][j], " ")
			}
			fmt.Fprintln(writer)
		}
	}
}

func solve(n int, m int, k int) [][]int {
	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, m)
	}

	var s int
	for pos := 0; pos < n*m; {
		for x := range k {
			r, c := pos/m, pos%m
			res[r][c] = (s+x)%k + 1
			pos++
		}
		if pos == n*m {
			break
		}
		for {
			r, c := pos/m, pos%m
			if (r == 0 || res[r-1][c] != s%k+1) && (c == 0 || res[r][c-1] != s%k+1) {
				break
			}
			s++
		}
	}

	return res
}
