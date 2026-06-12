package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	for {
		if _, err := fmt.Fscan(reader, &n); err != nil {
			break
		}
		moves := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &moves[i])
		}
		fmt.Fprintln(writer, solve(moves))
	}
}

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	moves := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &moves[i])
	}
	return solve(moves)
}

// solve counts the number of distinct possible headquarters locations.
//
// Rotate the grid 45 degrees with u = x+y, v = x-y. In (u,v) space each move
// becomes: U=(+1,-1), D=(-1,+1), L=(-1,-1), R=(+1,+1). For every record the two
// allowed moves either fix one coordinate and leave the other as +/-1:
//   - UL, DR  -> free in u (the u-diagonal), fixed in v
//   - UR, DL  -> free in v (the v-diagonal), fixed in u
//   - ULDR    -> free in both
//
// The two diagonals are independent, and a free moves over one diagonal yield
// a+1 distinct sums; likewise b+1 for the other. So the answer is (a+1)(b+1).
func solve(move []string) int {
	var a, b int
	for _, cur := range move {
		switch cur {
		case "UL", "DR":
			a++
		case "UR", "DL":
			b++
		case "ULDR":
			a++
			b++
		}
	}
	return (a + 1) * (b + 1)
}
