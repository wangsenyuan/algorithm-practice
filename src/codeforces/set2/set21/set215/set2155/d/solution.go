package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)

	ask := func(u int, v int) int {
		fmt.Println(u, v)
		var res int
		fmt.Fscan(reader, &res)
		return res
	}

	for range tc {
		var n int
		fmt.Fscan(reader, &n)
		solve(n, ask)
	}
}

func solve(n int, ask func(int, int) int) {
	for i := 1; i < n; i++ {
		for j := 1; j <= n; j++ {
			x := (i + j) % n
			if x == 0 {
				x = n
			}
			res := ask(j, x)
			if res == 1 {
				return
			}
		}
	}
}
