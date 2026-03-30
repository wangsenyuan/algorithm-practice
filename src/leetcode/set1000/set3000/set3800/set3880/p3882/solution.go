package p3882

import "math/big"

func minCost(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	// m * n <= 1000
	dist := make([][]*big.Int, n)
	for i := range n {
		dist[i] = make([]*big.Int, m)
		for j := range m {
			dist[i][j] = big.NewInt(0)
		}
	}

	dist[0][0].SetBit(dist[0][0], grid[0][0], 1)

	X := 1 << 10

	play := func(src, dst *big.Int, v int) {
	
		for x := range X {
			if src.Bit(x) == 1 {
				dst.SetBit(dst, x^v, 1)
			}
		}
	}

	for i := range n {
		for j := range m {
			if i > 0 {
				play(dist[i-1][j], dist[i][j], grid[i][j])
			}
			if j > 0 {
				play(dist[i][j-1], dist[i][j], grid[i][j])
			}
		}
	}

	for i := range X {
		if dist[n-1][m-1].Bit(i) == 1 {
			return i
		}
	}
	return -1
}
