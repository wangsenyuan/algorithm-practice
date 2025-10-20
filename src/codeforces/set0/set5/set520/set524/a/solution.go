package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, cur := range res {
		fmt.Fprintf(writer, "%d: %d", cur[0], len(cur)-1)
		for i := 1; i < len(cur); i++ {
			fmt.Fprintf(writer, " %d", cur[i])
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) [][]int {
	var m, k int
	fmt.Fscan(reader, &m, &k)
	friends := make([][]int, m)
	for i := range m {
		friends[i] = make([]int, 2)
		fmt.Fscan(reader, &friends[i][0], &friends[i][1])
	}
	return solve(k, friends)
}

func solve(k int, friends [][]int) [][]int {

	var users []int
	for _, cur := range friends {
		users = append(users, cur[0], cur[1])
	}

	slices.Sort(users)
	users = slices.Compact(users)

	n := len(users)
	adj := make([]map[int]bool, n)

	for i := range n {
		adj[i] = make(map[int]bool)
	}

	for _, cur := range friends {
		a, b := cur[0], cur[1]
		i := sort.SearchInts(users, a)
		j := sort.SearchInts(users, b)
		adj[i][j] = true
		adj[j][i] = true
	}

	res := make([][]int, n)
	for i := range n {
		// 自己的id是第一个
		res[i] = append(res[i], users[i])
	}

	play := func(x int, y int) {
		if x == y || adj[x][y] {
			return
		}
		var cnt int
		for c := range n {
			if adj[x][c] && adj[y][c] {
				cnt++
			}
		}
		if cnt*100 >= k*len(adj[x]) {
			res[x] = append(res[x], users[y])
		}
	}

	for i := range n {
		for j := range n {
			play(i, j)
		}
	}

	return res
}
