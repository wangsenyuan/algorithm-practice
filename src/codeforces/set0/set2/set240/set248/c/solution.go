package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res < 0 {
		fmt.Println(-1)
	} else {
		fmt.Printf("%.10f\n", res)
	}
}

func drive(reader *bufio.Reader) float64 {
	var y1, y2, yw, xb, yb, r int
	fmt.Fscan(reader, &y1, &y2, &yw, &xb, &yb, &r)
	return solve(y1, y2, yw, xb, yb, r)
}

func solve(y1, y2, yw, xb, yb, r int) float64 {
	wallY := yw - r
	goalY := y1 + r
	reflectedY := 2*wallY - yb

	// The distance from the upper post (0, y2) to the unfolded
	// trajectory must be at least r. Compare the squares exactly.
	delta := y2 - goalY
	dy := reflectedY - goalY
	left := square(int64(xb) * int64(delta))
	right := square(int64(r) * int64(xb))
	right.Add(right, square(int64(r)*int64(dy)))
	if left.Cmp(right) < 0 {
		return -1
	}

	return float64(xb) * float64(wallY-goalY) / float64(dy)
}

func square(v int64) *big.Int {
	x := big.NewInt(v)
	return x.Mul(x, x)
}
