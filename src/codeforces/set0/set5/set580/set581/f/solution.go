package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf int = 1e9

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	if n == 2 {
		return 1
	}
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Root at a non-leaf.
	var root int
	for i := range n {
		if len(adj[i]) > 1 {
			root = i
			break
		}
	}

	children := make([][]int, n)
	leafCount := make([]int, n)

	var build func(p, u int)

	build = func(p, u int) {
		if len(adj[u]) == 1 {
			leafCount[u] = 1
		}
		for _, v := range adj[u] {
			if v != p {
				children[u] = append(children[u], v)
				build(u, v)
				leafCount[u] += leafCount[v]
			}
		}
	}
	build(-1, root)

	totalLeaves := leafCount[root]
	half := totalLeaves / 2

	// z1[u][cnt][col] = min cross-edges in subtree of u, u has color col, exactly cnt leaves with color 0.
	// Stored as z1[u] = [leafCount[u]+1][2]int, index (cnt, col).
	z1 := make([][][]int, n)
	for i := range n {
		L := leafCount[i] + 1
		z1[i] = make([][]int, L)
		for c := range L {
			z1[i][c] = []int{inf, inf}
		}
	}

	var dfs func(u int)
	dfs = func(u int) {
		if len(children[u]) == 0 {
			z1[u][0][1] = 0 // leaf color 1, 0 leaves with color 0
			z1[u][1][0] = 0 // leaf color 0, 1 leaf with color 0
			return
		}
		ch := children[u]
		// Merge children from last to first; z2[cnt][col] = state after merging a suffix of children.
		// Initially: no children merged, so cnt=0 only.
		z2 := make([][]int, leafCount[u]+1)
		for i := range len(z2) {
			z2[i] = []int{inf, inf}
		}
		z2[0][0] = 0
		z2[0][1] = 0

		for idx := len(ch) - 1; idx >= 0; idx-- {
			s := ch[idx]
			dfs(s)
			Ls := leafCount[s]
			// newZ2[cnt][col] = min over a, ncol: z2[cnt-a][col] + z1[s][a][ncol] + (1 if ncol != col else 0)
			newZ2 := make([][]int, leafCount[u]+1)
			for i := range len(newZ2) {
				newZ2[i] = []int{inf, inf}
			}
			for cnt := 0; cnt <= leafCount[u]; cnt++ {
				for col := range 2 {
					best := inf
					for a := 0; a <= Ls && a <= cnt; a++ {
						prev := cnt - a
						if prev < 0 || prev >= len(z2) {
							continue
						}
						for ncol := range 2 {
							edgeCost := 0
							if ncol != col {
								edgeCost = 1
							}
							cur := z2[prev][col] + z1[s][a][ncol] + edgeCost
							if cur < best {
								best = cur
							}
						}
					}
					if best < inf {
						newZ2[cnt][col] = best
					}
				}
			}
			z2 = newZ2
		}
		for cnt := range len(z1[u]) {
			for col := range 2 {
				z1[u][cnt][col] = z2[cnt][col]
			}
		}
	}
	dfs(root)

	return min(z1[root][half][0], z1[root][half][1])
}
