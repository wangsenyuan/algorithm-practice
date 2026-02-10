package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, R, r int
	if _, err := fmt.Fscan(in, &n, &R, &r); err != nil {
		return
	}

	if solve(n, R, r) {
		fmt.Fprint(out, "YES")
	} else {
		fmt.Fprint(out, "NO")
	}
}

func solve(n int, R int, r int) bool {
	if r > R {
		return false
	}
	if n == 1 {
		return true
	}
	d := float64(R - r)
	if d <= 0 {
		return false
	}
	need := d * math.Sin(math.Pi/float64(n))
	return need+1e-9 >= float64(r)
}