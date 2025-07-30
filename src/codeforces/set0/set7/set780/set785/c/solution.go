package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	res := solve(n, m)
	fmt.Println(res)
}

func solve(n int, m int) int {
	if m >= n {
		return n
	}
	// m < n
	s := n - m

	i := sort.Search(s, func(i int) bool {
		if i == 0 {
			return false
		}
		return i+1 >= s*2/i
	})

	if i*(i+1)/2 < n-m {
		i++
	}

	return m + i
}
