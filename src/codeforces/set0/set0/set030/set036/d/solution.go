package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	defer writer.Flush()

	var t, k int
	fmt.Fscan(reader, &t, &k)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		if solve(n, m, k) {
			fmt.Fprintln(writer, "+")
		} else {
			fmt.Fprintln(writer, "-")
		}
	}
}

func solve(n, m, k int) bool {
	// Convert to 0-indexed: start at (0,0), target is (n-1, m-1)
	n--
	m--

	// Base case: already at terminal (can't move)
	if n < 0 || m < 0 {
		return false
	}

	if k == 1 {
		// Special case for k=1
		// Losing positions are when both n and m are even
		return !((n&1 == 0) && (m&1 == 0))
	}

	minVal := min(n, m)
	cycle := 2*k + 2
	rem := minVal % cycle

	if rem == k {
		// Stripe of 1s -> Winning
		return true
	}
	if rem == cycle-1 {
		// Stripe of 1s -> Winning
		// 2*k + 1
		return true
	}

	sum := n + m
	isOdd := (sum & 1) == 1

	if rem < k {
		// First stripe: alternating 0/1
		// Losing if sum is even
		return isOdd
	}

	// rem > k (since rem != k and rem != 2k+1)
	// Third stripe: inverted alternating
	// Losing if sum is odd
	return !isOdd
}
