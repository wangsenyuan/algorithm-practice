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

func drive(reader *bufio.Reader) int64 {
	var n int64
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int64) int64 {
	i := sort.Search(1e9, func(x int) bool {
		return int64(x*x) > n/2
	})
	i--
	j := sort.Search(1e9, func(x int) bool {
		return int64(x*x) > n/4
	})
	j--

	return int64(i + j)
}
