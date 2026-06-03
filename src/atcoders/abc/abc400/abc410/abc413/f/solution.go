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
	var h, w, n int
	fmt.Fscan(reader, &h, &w, &n)
	goals := make([][]int, n)
	for i := range n {
		goals[i] = make([]int, 2)
		fmt.Fscan(reader, &goals[i][0], &goals[i][1])
	}
	return solve(h, w, goals)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(h int, w int, goals [][]int) int {
	dist := make([][]int, h)
	mark := make([][]int, h)
	for i := range dist {
		dist[i] = make([]int, w)
		mark[i] = make([]int, w)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	que := make([]int, h*w)
	var head, tail int
	for _, cur := range goals {
		r, c := cur[0]-1, cur[1]-1
		dist[r][c] = 0
		que[head] = r*w + c
		head++
	}

	for tail < head {
		r, c := que[tail]/w, que[tail]%w
		tail++
		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < h && nc >= 0 && nc < w && dist[nr][nc] == -1 {
				mark[nr][nc]++
				if mark[nr][nc] == 2 {
					dist[nr][nc] = dist[r][c] + 1
					que[head] = nr*w + nc
					head++
				}
			}
		}
	}
	var res int
	for i := range h {
		for j := range w {
			res += max(0, dist[i][j])
		}
	}
	return res
}
