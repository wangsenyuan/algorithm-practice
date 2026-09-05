package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res, _, _ := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) ([]int, int, [][3]int) {
	var n, k, q int
	fmt.Fscan(reader, &n, &k, &q)
	queries := make([][3]int, q)
	for i := range queries {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(n, k, queries), k, queries
}

func solve(n, k int, queries [][3]int) []int {
	mark := make([]int, n+1)
	mark2 := make([]int, n+1)
	for _, cur := range queries {
		l, r := cur[1]-1, cur[2]-1

		if cur[0] == 1 {
			mark[l]++
			mark[r+1]--
		} else {
			mark2[l]++
			mark2[r+1]--
		}
	}
	for i := 1; i < n; i++ {
		mark[i] += mark[i-1]
		mark2[i] += mark2[i-1]
	}

	var arr []int
	res := make([]int, n)
	for i := range n {
		if mark[i] > 0 {
			res[i] = k + 1
			if mark2[i] == 0 {
				res[i] = k
			}
		} else if mark2[i] > 0 {
			arr = append(arr, i)
		}
	}

	for i, j := range arr {
		res[j] = i % k
	}

	return res
}
