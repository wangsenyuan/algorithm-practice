package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	var a string
	fmt.Fscan(reader, &n, &a)
	return solve(a)
}

func solve(a string) []int {
	n := len(a)
	ids := make([]int, n)
	for i := range n {
		ids[i] = i
	}
	next := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		next[i] = i + 1
		if i+1 < n && a[i] == a[i+1] {
			next[i] = next[i+1]
		}
	}

	// aba
	// ba
	// aa
	cmp := func(i int, j int) int {
		// i < j
		i1 := next[i]
		if i1 == n {
			return -1
		}
		// i1 < n
		// i1 是第一个和i不同的字符
		if i1 <= j {
			// [i...i1)都是一样的
			if a[i1] < a[i] {
				return -1
			}
			return 1
		}
		// i1 > j
		// aaaaaabc
		// aaaaaabc
		return -1
	}

	slices.SortFunc(ids, func(i int, j int) int {
		if i < j {
			return cmp(i, j)
		}
		return -cmp(j, i)
	})

	for i := range n {
		ids[i]++
	}

	return ids
}
