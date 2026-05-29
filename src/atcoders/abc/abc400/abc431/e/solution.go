package main

import (
	"bufio"
	"container/heap"
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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []string) int {
	n := len(a)
	m := len(a[0])
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	dp := make([][][4]int, n)
	for i := range n {
		dp[i] = make([][4]int, m)
		for j := range m {
			for d := range 4 {
				dp[i][j][d] = 1 << 60
			}
		}
	}

	type state struct {
		x    int
		y    int
		dir  int
		cost int
	}

	var que queue[state]

	que.push(state{0, -1, 1, 0})

	for !que.isEmpty() {
		cur := que.pop()
		nx, ny := cur.x+dx[cur.dir], cur.y+dy[cur.dir]
		if nx < 0 || nx == n || ny < 0 || ny == m {
			continue
		}
		x := int(a[nx][ny] - 'A')
		if a[nx][ny] == 'C' {
			// 3
			x++
		}
		for ndir := range 4 {
			if cur.dir^ndir == x {
				if cur.cost < dp[nx][ny][ndir] {
					dp[nx][ny][ndir] = cur.cost
					que.pushLeft(state{nx, ny, ndir, cur.cost})
				}
			} else if cur.dir^ndir != 2 {
				if cur.cost+1 < dp[nx][ny][ndir] {
					dp[nx][ny][ndir] = cur.cost + 1
					que.push(state{nx, ny, ndir, cur.cost + 1})
				}
			}
		}
	}

	return dp[n-1][m-1][1]
}

type queue[T any] struct {
	right []T
	left  []T
}

func (que *queue[T]) push(x T) {
	que.right = append(que.right, x)
}

func (que *queue[T]) pushLeft(x T) {
	que.left = append(que.left, x)
}

func (que *queue[T]) isEmpty() bool {
	return len(que.left)+len(que.right) == 0
}

func (que *queue[T]) pop() T {
	if len(que.left) > 0 {
		n := len(que.left)
		last := que.left[n-1]
		que.left = que.left[:n-1]
		return last
	}
	first := que.right[0]
	que.right = que.right[1:]
	return first
}

func solve1(a []string) int {
	n := len(a)
	m := len(a[0])
	// dp[i][j][0/1/2/3]表示到(i,j), 且从上/右/下/左方向射入的最小修改数
	dp := make([][][4]int, n)
	for i := range n {
		dp[i] = make([][4]int, m)
		for j := range m {
			for d := range 4 {
				dp[i][j][d] = 1 << 60
			}
		}
	}

	var pq PQ

	add := func(r int, c int, d int, v int) {
		if r >= 0 && r < n && c >= 0 && c < m && v < dp[r][c][d] {
			dp[r][c][d] = v
			heap.Push(&pq, node{r, c, d, v})
		}
	}

	var dd = [][]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}
	mirro_a := []int{0, 1, 2, 3}
	mirro_b := []int{3, 2, 1, 0}
	mirro_c := []int{1, 0, 3, 2}

	changes := [][]int{
		mirro_a, mirro_b, mirro_c,
	}

	add(0, 0, 3, 0)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(node)
		r, c, d := cur.r, cur.c, cur.d
		if cur.priority != dp[r][c][d] {
			continue
		}
		st := int(a[r][c] - 'A')

		for i := range 3 {
			pos := (st + i) % 3
			nd := changes[pos][d]
			nr, nc := r+dd[nd][0], c+dd[nd][1]
			add(nr, nc, nd, dp[r][c][d]+checkBool(i > 0))
		}
	}

	ans := 1 << 60
	for d := range 4 {
		if dp[n-1][m-1][d] < 0 {
			// 会出现这个状态吗？好像会的
			continue
		}
		st := int(a[n-1][m-1] - 'A')
		for i := range 3 {
			pos := (st + i) % 3
			nd := changes[pos][d]
			if nd == 3 {
				// 只能从右边出去
				ans = min(ans, dp[n-1][m-1][d]+checkBool(i > 0))
			}
		}
	}

	return ans
}

func checkBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

type node struct {
	r        int
	c        int
	d        int
	priority int
}

type PQ []node

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i int, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(node))
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}
