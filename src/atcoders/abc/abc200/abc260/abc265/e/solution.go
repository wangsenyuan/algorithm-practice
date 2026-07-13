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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	moves := make([][]int, 3)
	for i := range 3 {
		moves[i] = make([]int, 2)
		fmt.Fscan(reader, &moves[i][0], &moves[i][1])
	}
	obs := make([][]int, m)
	for i := range m {
		obs[i] = make([]int, 2)
		fmt.Fscan(reader, &obs[i][0], &obs[i][1])
	}
	return solve(n, moves, obs)
}

type pair struct {
	first  int
	second int
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(n int, moves [][]int, obs [][]int) int {
	dp := make([][]int, n+1)
	ndp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, n+1)
		ndp[i] = make([]int, n+1)
	}

	block := make(map[pair]bool)
	for _, cur := range obs {
		block[pair{cur[0], cur[1]}] = true
	}

	move := func(x0 int, y0 int, cur []int, cnt int) (x int, y int) {
		dx, dy := cur[0], cur[1]
		x = x0 + dx*cnt
		y = y0 + dy*cnt
		return
	}

	dp[0][0] = 1

	for w := range n {
		for x := range n {
			for y := range n {
				ndp[x][y] = 0
			}
		}
		// x + y + z = w
		for x := range w + 1 {
			for y := 0; x+y <= w; y++ {
				z := w - x - y
				r, c := move(0, 0, moves[0], x)
				r, c = move(r, c, moves[1], y)
				r, c = move(r, c, moves[2], z)
				nr, nc := move(r, c, moves[0], 1)
				if !block[pair{nr, nc}] {
					ndp[x+1][y] = add(ndp[x+1][y], dp[x][y])
				}
				nr, nc = move(r, c, moves[1], 1)
				if !block[pair{nr, nc}] {
					ndp[x][y+1] = add(ndp[x][y+1], dp[x][y])
				}
				nr, nc = move(r, c, moves[2], 1)
				if !block[pair{nr, nc}] {
					ndp[x][y] = add(ndp[x][y], dp[x][y])
				}
			}
		}
		dp, ndp = ndp, dp
	}

	var ans int
	for x := range n + 1 {
		for y := range n + 1 {
			ans = add(ans, dp[x][y])
		}
	}

	return ans
}
