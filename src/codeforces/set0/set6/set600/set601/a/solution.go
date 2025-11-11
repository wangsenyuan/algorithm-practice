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
	railways := make([][]int, m)
	for i := range m {
		railways[i] = make([]int, 2)
		fmt.Fscan(reader, &railways[i][0], &railways[i][1])
	}
	return solve(n, railways)
}

func solve(n int, railways [][]int) int {
	if len(railways) == 0 {
		return -1
	}
	direct_train := false
	for _, cur := range railways {
		u, v := cur[0], cur[1]
		u, v = min(u, v), max(u, v)
		if u == 1 && v == n {
			direct_train = true
			break
		}
	}

	dist := make([][]int, n)

	for i := range n {
		dist[i] = make([]int, n)
		for j := range n {
			dist[i][j] = inf
		}
		dist[i][i] = 0
	}

	if direct_train {
		for i := range n {
			for j := range n {
				if i != j {
					dist[i][j] = 1
				}
			}
		}
	}

	for _, cur := range railways {
		u, v := cur[0]-1, cur[1]-1
		if direct_train {
			// 这条路不能用来跑车
			dist[u][v] = inf
			dist[v][u] = inf
		} else {
			dist[u][v] = 1
			dist[v][u] = 1
		}
	}

	for k := range n {
		for i := range n {
			for j := range n {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	if dist[0][n-1] == inf {
		return -1
	}
	return dist[0][n-1]
}

const inf = 1 << 60
