package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	res := solve(n, m)
	fmt.Println(res)
}

func solve(n int, m int) int {
	n, m = max(n, m), min(n, m)
	if m == 1 {
		return n
	}
	if m == 2 {
		return 2 * (2*(n/4) + min(n%4, 2))
	}

	if m == 3 && (n == 3 || n == 5 || n == 6) || m == 4 && n == 4 {
		return bruteForce(m, n)
	}

	return m*n - (m*n)/2
}

func bruteForce(m int, n int) int { // m * n <= 18
	var best int
	checkAndSet := func(state int) {

		get := func(i int, j int) int {
			if i < 0 || i >= n || j < 0 || j >= m {
				return 0
			}
			pos := i*m + j
			return (state >> pos) & 1
		}

		cnt := bits.OnesCount(uint(state))
		if cnt <= best {
			return
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if get(i, j) == 1 && (get(i-1, j-2) == 1 || get(i-2, j-1) == 1 || get(i-1, j+2) == 1 || get(i-2, j+1) == 1) {
					return
				}
			}
		}
		best = cnt
	}

	for state := 1; state < 1<<(m*n); state++ {
		checkAndSet(state)
	}
	return best
}
