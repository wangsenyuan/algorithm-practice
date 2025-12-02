package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, x int
	fmt.Fscan(reader, &n, &m, &x)
	res := solve(n, m, x)
	fmt.Println(res)
}

func solve(n int, m int, x int) int {
	var res int

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if (i-1+j-1)%2 == 0 {
				// black
				w := min(i, j, n-i+1, m-j+1)
				if w == x {
					res++
				}
			}
		}
	}
	return res
}
