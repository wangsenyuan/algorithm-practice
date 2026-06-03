package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	var s, t string
	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &t)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}

	return solve(s, t, queries)
}

func solve(s string, t string, queries [][]int) string {
	n := len(s)
	diff := make([]int, n+1)
	for _, cur := range queries {
		l, r := cur[0], cur[1]
		diff[l-1]++
		diff[r]--
	}

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	buf := []byte(s)

	for i := 0; i < n; i++ {
		if diff[i]&1 == 1 {
			buf[i] = t[i]
		}
	}
	return string(buf)
}
