package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	if res == nil {
		fmt.Println(-1)
	} else {
		fmt.Println(res[0], res[1], res[2])
	}
}

func solve(n int) []int {
	if n == 1 {
		return nil
	}

	return []int{n, n + 1, n * (n + 1)}
}
