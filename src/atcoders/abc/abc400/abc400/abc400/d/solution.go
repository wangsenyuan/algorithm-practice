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
	var h, w int
	fmt.Fscan(reader, &h, &w)
	grid := make([]string, h)
	for i := range h {
		fmt.Fscan(reader, &grid[i])
	}
	var a, b, c, d int
	fmt.Fscan(reader, &a, &b, &c, &d)
	return solve(h, w, grid, a, b, c, d)
}

const inf = 1 << 60

var dd = []int{-1, 0, 1, 0, -1}

func solve(h, w int, grid []string, a, b, c, d int) int {
	a--
	b--
	c--
	d--

	dist := make([][]int, h)
	for i := range h {
		dist[i] = make([]int, w)
		for j := range w {
			dist[i][j] = inf
		}
	}

	var dq Dqueue[node]

	dq.PushTail(node{a, b, 0})
	dist[a][b] = 0

	for !dq.IsEmpty() {
		cur := dq.PopFront()
		r, c := cur.r, cur.c
		if dist[r][c] != cur.dist {
			continue
		}

		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr < 0 || nr == h || nc < 0 || nc == w {
				continue
			}
			if grid[nr][nc] == '.' && dist[nr][nc] > dist[r][c] {
				dist[nr][nc] = dist[r][c]
				dq.PushFront(node{nr, nc, dist[nr][nc]})
			}
		}
		// kick
		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < h && nc >= 0 && nc < w {
				if dist[nr][nc] > dist[r][c]+1 {
					dist[nr][nc] = dist[r][c] + 1
					dq.PushTail(node{nr, nc, dist[nr][nc]})
				}
				nr += dd[i]
				nc += dd[i+1]
				if nr >= 0 && nr < h && nc >= 0 && nc < w {
					if dist[nr][nc] > dist[r][c]+1 {
						dist[nr][nc] = dist[r][c] + 1
						dq.PushTail(node{nr, nc, dist[nr][nc]})
					}
				}
			}
		}
	}

	return dist[c][d]
}

type node struct {
	r    int
	c    int
	dist int
}

type Dqueue[T any] struct {
	front []T
	tail  []T
}

func (this *Dqueue[T]) IsEmpty() bool {
	return len(this.front) == 0 && len(this.tail) == 0
}

func (this *Dqueue[T]) PushFront(x T) {
	this.front = append(this.front, x)
}

func (this *Dqueue[T]) PushTail(x T) {
	this.tail = append(this.tail, x)
}

func (this *Dqueue[T]) PopFront() T {
	if len(this.front) > 0 {
		ret := this.front[len(this.front)-1]
		this.front = this.front[:len(this.front)-1]
		return ret
	}
	ret := this.tail[0]
	this.tail = this.tail[1:]
	return ret
}
