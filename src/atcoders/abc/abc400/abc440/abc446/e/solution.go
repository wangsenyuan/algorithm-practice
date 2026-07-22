package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var m, a, b int
	fmt.Fscan(reader, &m, &a, &b)
	return solve(m, a, b)
}

func solve(m, a, b int) int {
	adj := make([][]int, m*m)
	for x := 1; x < m; x++ {
		for y := 1; y < m; y++ {
			z := (a*y + b*x) % m
			adj[y*m+z] = append(adj[y*m+z], x*m+y)
		}
	}

	marked := make([]bool, m*m)
	var que []int
	for y := 1; y < m; y++ {
		marked[y*m] = true
		que = append(que, y*m+0)
	}

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		for _, next := range adj[cur] {
			if !marked[next] {
				marked[next] = true
				que = append(que, next)
			}
		}
	}

	var cnt int
	for x := 1; x < m; x++ {
		for y := 1; y < m; y++ {
			if !marked[x*m+y] {
				cnt++
			}
		}
	}

	return cnt
}
