package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, s int
	fmt.Fscan(reader, &n, &s)
	return solve(n, s)
}

func solve(n int, s int) int {
	check := func(x int) bool {
		var ds int
		for i := x; i > 0; i /= 10 {
			ds += i % 10
		}
		return x-ds >= s
	}
	if !check(n) {
		return 0
	}
	l := sort.Search(n, check)
	return n - l + 1
}
