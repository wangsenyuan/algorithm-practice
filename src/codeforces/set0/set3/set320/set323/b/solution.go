package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, row := range res {
		for _, val := range row {
			fmt.Fprint(writer, val, " ")
		}
		fmt.Fprintln(writer)
	}
}

func solve(n int) [][]int {
	// Only n=4 is impossible; for all other n>=3 we can construct such a tournament.
	// Odd n: regular cyclic tournament (each vertex points to next (n-1)/2 vertices) has diameter 2.
	if n%2 == 1 {
		res := make([][]int, n)
		for i := range n {
			res[i] = make([]int, n)
			for j := 1; j <= n/2; j++ {
				res[i][(i+j)%n] = 1
			}
		}
		return res
	}

	if n == 4 {
		return nil
	}

	// Even n >= 6:
	// Use a direct O(n^2) construction.
	//
	// Base: a known diameter-2 tournament for n=6.
	// Then for each step, add a 2-vertex gadget (a, b) in that order:
	// - a beats all previous vertices
	// - all previous vertices beat b
	// - b -> a
	//
	// This preserves tournament property, and ensures diameter <= 2:
	// - any old v reaches a via v -> b -> a (since v -> b, b -> a)
	// - b reaches any old v via b -> a -> v (since a -> v)
	// - a reaches b via a -> v -> b for any old v (since a -> v, v -> b)
	base6 := [][]int{
		{0, 1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 1},
		{1, 0, 1, 0, 0, 1},
		{1, 1, 0, 0, 0, 0},
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// copy base 6x6
	for i := 0; i < 6; i++ {
		copy(res[i][:6], base6[i])
	}

	// add gadgets without re-copying (still matches the iterative construction)
	for a := 6; a < n; a += 2 {
		b := a + 1
		// b -> a
		res[b][a] = 1
		// connect (a,b) with all previous vertices v < a
		for v := 0; v < a; v++ {
			res[a][v] = 1 // a beats v
			res[v][b] = 1 // v beats b
		}
	}

	return res
}
