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
		for _, line := range drive(reader) {
			fmt.Fprintln(writer, line)
		}
	}
}

func drive(reader *bufio.Reader) []string {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	var s string
	fmt.Fscan(reader, &s)
	queries := readQueries(reader, q)
	_ = n
	return solve(s, queries)
}

func readQueries(reader *bufio.Reader, q int) [][3]int {
	queries := make([][3]int, q)
	for i := range queries {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return queries
}

func solve(s string, queries [][3]int) []string {
	n := len(s)
	todo := make([][]int, n)
	for i, cur := range queries {
		r := cur[1] - 1
		todo[r] = append(todo[r], i)
	}

	flag := make([]int, len(queries))

	val := make([]int, n)

	pref := make([]int, n)

	play := func(pat string, d int) {
		// 如果 pat是 0101
		clear(pref)
		clear(val)
		var cnt int
		for i := range n {
			x := int(s[i] - '0')
			y := int(pat[i&1] - '0')
			if x != y {
				val[i] = 1
				// 这里必须要翻转一次
				if i == 0 || val[i-1] == 0 {
					cnt++
				}
			}
			pref[i] = cnt
			for _, j := range todo[i] {
				l, k := queries[j][0]-1, queries[j][2]
				k1 := cnt
				if l > 0 {
					k1 -= pref[l-1]
				}
				if val[l] > 0 && l > 0 && val[l-1] == 1 {
					k1++
				}
				if k1 <= k {
					flag[j] |= d
				}
			}
		}
	}
	play("01", 1)
	play("10", 2)

	ans := make([]string, len(queries))

	for i := range flag {
		if flag[i] > 0 {
			ans[i] = "YES"
		} else {
			ans[i] = "NO"
		}
	}

	return ans
}
