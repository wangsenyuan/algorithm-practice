package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (n int, load [][]int, res []int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	load = make([][]int, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		load[i] = []int{a, b}
	}
	res = solve(n, load)
	return
}

func solve(n int, loan [][]int) []int {
	adj := make([]map[int]bool, n+1)
	for i := range n + 1 {
		adj[i] = make(map[int]bool)
	}
	for _, cur := range loan {
		a, b := cur[0], cur[1]
		adj[a][b] = true
	}
	res := make([]int, n)
	for u := 0; u < n; u++ {
		found := false
		for i := u - 1; i >= 0; i-- {
			v := res[i]
			if !adj[v][u+1] {
				// v欠u的钱，所以v不能在u的前面
				res[i+1] = u + 1
				found = true
				break
			}
			res[i+1] = v
		}
		if !found {
			res[0] = u + 1
		}
	}
	return res
}
