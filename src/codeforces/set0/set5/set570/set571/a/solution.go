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
	var a, b, c, l int
	fmt.Fscan(reader, &a, &b, &c, &l)
	return solve(a, b, c, l)
}

func solve(a int, b int, c int, l int) int {
	total := countTotalWays(l)
	// Count invalid cases where each stick is >= sum of other two
	// These cases are mutually exclusive (at most one can be true for positive values)
	invalid := countInvalidCase(a, b, c, l) +
		countInvalidCase(b, a, c, l) +
		countInvalidCase(c, a, b, l)
	return total - invalid
}

// countTotalWays counts all ways to distribute at most l among 3 sticks
// For each x (0 <= x <= l), ways to distribute x among 3 sticks = C(x+2, 2) = (x+2)(x+1)/2
func countTotalWays(l int) int {
	total := 0
	for x := 0; x <= l; x++ {
		total += (x + 2) * (x + 1) / 2
	}
	return total
}

// countInvalidCase counts cases where first stick >= sum of other two
// Assumes we're checking: a' >= b' + c'
// where a' = a + x, b' = b + y, c' = c + z
// We need: a + x >= (b + y) + (c + z)
// Which means: x >= (b + c - a) + (y + z)
// Let d = b + c - a (can be negative)
func countInvalidCase(a int, b int, c int, l int) int {
	d := b + c - a // can be negative

	count := 0
	// Iterate over all possible values of s = y + z (0 <= s <= l)
	for s := 0; s <= l; s++ {
		// Number of ways to get y + z = s: (s + 1) ways
		// (y can be 0, 1, 2, ..., s, and z = s - y)
		waysToGetS := s + 1

		// x must satisfy: x >= d + s (from triangle inequality failure: a+x >= b+y+c+z)
		// and x + s <= l (total increase constraint: x + y + z <= l)
		// and x >= 0 (non-negative increase)
		// So x is in range [max(0, d + s), l - s]
		xMin := max(0, d+s)
		xMax := l - s

		if xMin <= xMax {
			validX := xMax - xMin + 1
			count += waysToGetS * validX
		}
	}

	return count
}
