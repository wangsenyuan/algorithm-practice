package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return solve(n, edges)
}

const limit = int64(1e18)
const levels = 60

func solve(n int, edges [][]int) int {
	words := (n + 63) >> 6

	newMatrix := func() [][]uint64 {
		mat := make([][]uint64, n)
		for i := 0; i < n; i++ {
			mat[i] = make([]uint64, words)
		}
		return mat
	}

	P := make([][][]uint64, levels)
	Q := make([][][]uint64, levels)
	P[0] = newMatrix()
	Q[0] = newMatrix()

	for _, e := range edges {
		u, v, t := e[0]-1, e[1]-1, e[2]
		if t == 0 {
			P[0][u][v>>6] |= 1 << (v & 63)
		} else {
			Q[0][u][v>>6] |= 1 << (v & 63)
		}
	}

	mul := func(a [][]uint64, b [][]uint64) [][]uint64 {
		res := newMatrix()
		for i := 0; i < n; i++ {
			for block, mask := range a[i] {
				cur := mask
				for cur != 0 {
					lsb := cur & -cur
					bit := bits.TrailingZeros64(cur)
					u := block*64 + bit
					if u < n {
						for w := 0; w < words; w++ {
							res[i][w] |= b[u][w]
						}
					}
					cur ^= lsb
				}
			}
		}
		return res
	}

	for k := 1; k < levels; k++ {
		P[k] = mul(P[k-1], Q[k-1])
		Q[k] = mul(Q[k-1], P[k-1])
	}

	advance := func(cur []uint64, mat [][]uint64) []uint64 {
		next := make([]uint64, words)
		for block, mask := range cur {
			x := mask
			for x != 0 {
				lsb := x & -x
				bit := bits.TrailingZeros64(x)
				u := block*64 + bit
				if u < n {
					for w := 0; w < words; w++ {
						next[w] |= mat[u][w]
					}
				}
				x ^= lsb
			}
		}
		return next
	}

	empty := func(bs []uint64) bool {
		for _, x := range bs {
			if x != 0 {
				return false
			}
		}
		return true
	}

	cur := make([]uint64, words)
	cur[0] = 1 // vertex 1
	var ans int64

	for k := levels - 1; k >= 0; k-- {
		var next []uint64
		if bits.OnesCount64(uint64(ans))&1 == 0 {
			next = advance(cur, P[k])
		} else {
			next = advance(cur, Q[k])
		}
		if empty(next) {
			continue
		}
		ans += 1 << k
		cur = next
		if ans > limit {
			return -1
		}
	}

	return int(ans)
}
