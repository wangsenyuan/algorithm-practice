package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var k int
	fmt.Fscan(reader, &k)
	fmt.Println(solve(k))
}

func solve(k int) string {
	if k == 0 {
		return "a"
	}
	if k == 1 {
		return "aa"
	}

	find := func(k int) int {
		// (n + 1) * n / 2 <= k
		return sort.Search(k+2, func(n int) bool {
			return (n-1)*n/2 > k
		}) - 1
	}

	var buf bytes.Buffer
	letter := byte('a')
	for k > 0 {
		n := find(k)
		for range n {
			buf.WriteByte(letter)
		}
		k -= n * (n - 1) / 2
		letter++
	}

	return buf.String()
}
