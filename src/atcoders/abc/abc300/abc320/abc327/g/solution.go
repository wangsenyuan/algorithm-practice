package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod int64 = 998244353

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	return solve(n, m)
}

func solve(n, m int) int {
	limit := n * n / 4
	maxValue := n * n
	comb := make([][]int64, maxValue+1)
	for i := range comb {
		comb[i] = make([]int64, maxValue+1)
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = comb[i-1][j-1]
			if j < i {
				comb[i][j] += comb[i-1][j]
			}
			comb[i][j] %= mod
		}
	}

	maxEdges := make([]int, n+1)
	for i := 0; i <= n; i++ {
		maxEdges[i] = i * i / 4
	}

	g := make([][]int64, n+1)
	h := make([][]int64, n+1)
	f := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		g[i] = make([]int64, limit+1)
		h[i] = make([]int64, limit+1)
		f[i] = make([]int64, limit+1)
		for black := 0; black <= i; black++ {
			edges := black * (i - black)
			for e := 0; e <= edges; e++ {
				g[i][e] = (g[i][e] + comb[i][black]*comb[edges][e]) % mod
			}
		}
	}

	// h[v][e] counts connected, properly two-colored simple graphs.
	for vertices := 1; vertices <= n; vertices++ {
		for edges := 0; edges <= maxEdges[vertices]; edges++ {
			value := g[vertices][edges]
			for firstSize := 1; firstSize < vertices; firstSize++ {
				ways := comb[vertices-1][firstSize-1]
				for firstEdges := 0; firstEdges <= edges; firstEdges++ {
					if firstEdges > maxEdges[firstSize] || edges-firstEdges > maxEdges[vertices-firstSize] {
						continue
					}
					value -= ways * h[firstSize][firstEdges] % mod * g[vertices-firstSize][edges-firstEdges] % mod
				}
			}
			h[vertices][edges] = (value%mod + mod) % mod
		}
	}

	// A connected bipartite graph has exactly two proper two-colorings.
	inv2 := (mod + 1) / 2
	f[0][0] = 1
	for vertices := 1; vertices <= n; vertices++ {
		for componentSize := 1; componentSize <= vertices; componentSize++ {
			ways := comb[vertices-1][componentSize-1]
			for componentEdges := 0; componentEdges <= maxEdges[componentSize]; componentEdges++ {
				connected := h[componentSize][componentEdges] * inv2 % mod
				if connected == 0 {
					continue
				}
				for restEdges := 0; restEdges <= maxEdges[vertices-componentSize]; restEdges++ {
					if componentEdges+restEdges > maxEdges[vertices] {
						continue
					}
					f[vertices][componentEdges+restEdges] += ways * connected % mod * f[vertices-componentSize][restEdges] % mod
					f[vertices][componentEdges+restEdges] %= mod
				}
			}
		}
	}

	// Replace each of k simple edges with a non-empty set of labeled parallel edges.
	balls := make([]int64, limit+1)
	for boxes := 1; boxes <= limit; boxes++ {
		for used := 1; used <= boxes; used++ {
			term := comb[boxes][used] * modPow(int64(used), m) % mod
			if (boxes-used)&1 == 1 {
				balls[boxes] -= term
			} else {
				balls[boxes] += term
			}
		}
		balls[boxes] = (balls[boxes]%mod + mod) % mod
	}

	var ans int64
	for edges := 1; edges <= limit; edges++ {
		ans += f[n][edges] * balls[edges] % mod
		ans %= mod
	}
	// Each undirected edge has two orientations: (u, v) or (v, u).
	ans = ans * modPow(2, m) % mod
	return int(ans)
}

func modPow(a int64, b int) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}
