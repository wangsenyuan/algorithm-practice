package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int) int {
	if n == 0 {
		return 0
	}

	check := func(a, b int) bool {
		return (a+b)*(a*a+b*b) >= n
	}

	res := 1_000_000_000_000_000_000
	for a := 0; a <= 1e6; a++ {
		b := sort.Search(1_000_001, func(b int) bool {
			return check(a, b)
		})
		if b < a {
			break
		}
		w := a*a*a + a*a*b + a*b*b + b*b*b
		if w >= n {
			res = min(res, w)
		}
	}

	return res
}
