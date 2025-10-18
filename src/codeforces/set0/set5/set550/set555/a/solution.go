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

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	chains := make([][]int, m)
	for i := range m {
		var k int
		fmt.Fscan(reader, &k)
		chains[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &chains[i][j])
		}
	}
	return solve(chains, n)
}

func solve(chains [][]int, n int) int {
	at_chain := make([]int, n+1)
	pos := make([]int, n+1)
	for i := range n + 1 {
		at_chain[i] = -1
		pos[i] = -1
	}
	for i, cur := range chains {
		for j, v := range cur {
			at_chain[v-1] = i
			pos[v-1] = j
		}
	}

	// 如果i是独立的，那么只需要一次把它连接到chain里面去

	var ans int

	v := 1
	for v < n && at_chain[v] == at_chain[0] {
		v++
	}

	for v < n {
		if at_chain[v] != at_chain[0] && pos[v] == len(chains[at_chain[v]])-1 {
			// 这类需要一次操作
			ans++
		} else {
			ans += 2
		}

		v++
	}

	return ans
}
