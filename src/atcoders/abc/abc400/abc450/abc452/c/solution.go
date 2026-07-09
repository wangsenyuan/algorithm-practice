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

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []string {
	var n int
	fmt.Fscan(reader, &n)
	ribs := make([][]int, n)
	for i := range n {
		ribs[i] = make([]int, 2)
		fmt.Fscan(reader, &ribs[i][0], &ribs[i][1])
	}
	var m int
	fmt.Fscan(reader, &m)
	words := make([]string, m)
	for i := range m {
		fmt.Fscan(reader, &words[i])
	}
	return solve(ribs, words)
}

func solve(ribs [][]int, words []string) []string {
	m := len(words)

	var adj [26][11][11]int

	for i := range m {
		k := len(words[i])
		for j := range k {
			x := int(words[i][j] - 'a')
			adj[x][k][j]++
		}
	}

	n := len(ribs)
	// 可以重复使用, 所以不是一个

	// len(rib[i]) = A[i]
	// ribs[i][B[i]] = i
	check := func(strip string) bool {
		if len(strip) != n {
			return false
		}

		for i := range n {
			k := ribs[i][0]
			p := ribs[i][1]
			x := int(strip[i] - 'a')
			if adj[x][k][p-1] == 0 {
				return false
			}
		}

		return true
	}

	ans := make([]string, m)

	for i := range m {
		if check(words[i]) {
			ans[i] = "Yes"
		} else {
			ans[i] = "No"
		}
	}

	return ans
}
