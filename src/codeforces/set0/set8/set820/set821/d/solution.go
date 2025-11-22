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
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	lambs := make([][]int, k)
	for i := range k {
		lambs[i] = make([]int, 2)
		fmt.Fscan(reader, &lambs[i][0], &lambs[i][1])
	}
	return solve(n, m, lambs)
}

type pair struct {
	first  int
	second int
}

func solve(n int, m int, lambs [][]int) int {
	k := len(lambs)
	T := n + m + k
	adj := make([][]pair, T)

	labels := make(map[pair]int)
	rowCells := make(map[int][]int)
	colCells := make(map[int][]int)

	for i := range k {
		r, c := lambs[i][0]-1, lambs[i][1]-1
		labels[pair{r, c}] = i
		rowCells[r] = append(rowCells[r], i)
		colCells[c] = append(colCells[c], i)
	}

	// Build edges from lit cells to rows/columns and between adjacent lit cells
	for i := range k {
		r, c := lambs[i][0]-1, lambs[i][1]-1
		// From lit cell to its row and adjacent rows (cost 1 to light the row)
		for dr := -1; dr <= 1; dr++ {
			nr := r + dr
			if nr >= 0 && nr < n {
				adj[i] = append(adj[i], pair{k + nr, 1})
			}
		}
		// From lit cell to its column and adjacent columns (cost 1 to light the column)
		for dc := -1; dc <= 1; dc++ {
			nc := c + dc
			if nc >= 0 && nc < m {
				adj[i] = append(adj[i], pair{k + n + nc, 1})
			}
		}

		// Adjacent lit cells (4-directional movement, cost 0)
		for _, dx := range []int{-1, 0, 1} {
			for _, dy := range []int{-1, 0, 1} {
				if (dx == 0 && dy == 0) || (dx != 0 && dy != 0) {
					continue
				}
				if j, ok := labels[pair{r + dx, c + dy}]; ok && j != i {
					adj[i] = append(adj[i], pair{j, 0})
				}
			}
		}
	}

	// Build edges from rows to all lit cells in those rows and adjacent rows (cost 0)
	for r := range n {
		for dr := -1; dr <= 1; dr++ {
			nr := r + dr
			if nr >= 0 && nr < n {
				if cells, ok := rowCells[nr]; ok {
					for _, j := range cells {
						adj[k+r] = append(adj[k+r], pair{j, 0})
					}
				}
			}
		}
	}

	// Build edges from columns to all lit cells in those columns and adjacent columns (cost 0)
	for c := range m {
		for dc := -1; dc <= 1; dc++ {
			nc := c + dc
			if nc >= 0 && nc < m {
				if cells, ok := colCells[nc]; ok {
					for _, j := range cells {
						adj[k+n+c] = append(adj[k+n+c], pair{j, 0})
					}
				}
			}
		}
	}

	dist := make([]int, T)
	for i := range T {
		dist[i] = inf
	}

	que := make([]int, 2*T)
	head, tail := T, T
	que[head] = labels[pair{0, 0}]
	head++
	dist[labels[pair{0, 0}]] = 0

	for tail < head {
		u := que[tail]
		tail++
		for _, cur := range adj[u] {
			v := cur.first
			if cur.second == 1 && dist[v] > dist[u]+1 {
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			} else if cur.second == 0 && dist[v] > dist[u] {
				dist[v] = dist[u]
				tail--
				que[tail] = v
			}
		}
	}

	ans := inf

	// Check if destination (n-1, m-1) is a lit cell
	if destIdx, ok := labels[pair{n - 1, m - 1}]; ok {
		ans = min(ans, dist[destIdx])
	}

	// Check if we can reach destination by lighting the last row or column from a lit cell
	for i := range k {
		r, c := lambs[i][0]-1, lambs[i][1]-1
		// If we're at a lit cell in the last row, we can light that row (cost 1) to reach (n-1, m-1)
		if r == n-1 {
			ans = min(ans, dist[i]+1)
		}
		// If we're at a lit cell in the last column, we can light that column (cost 1) to reach (n-1, m-1)
		if c == m-1 {
			ans = min(ans, dist[i]+1)
		}
	}

	// If we've reached the last row node, we can reach (n-1, m-1) which is in that row
	// The cost to light the row is already included in dist[k+n-1]
	if dist[k+n-1] < inf {
		ans = min(ans, dist[k+n-1])
	}
	// If we've reached the last column node, we can reach (n-1, m-1) which is in that column
	// The cost to light the column is already included in dist[k+n+m-1]
	if dist[k+n+m-1] < inf {
		ans = min(ans, dist[k+n+m-1])
	}

	// Also check if we can reach destination by being at a lit cell near the destination
	// and lighting both the row and column (but that would cost 2, which is already covered above)

	if ans == inf {
		return -1
	}
	return ans
}

const inf = 1 << 60
