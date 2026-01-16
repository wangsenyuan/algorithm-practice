package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	res := process(reader)
	fmt.Fprintf(w, "%.8f\n", res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) float64 {
	n := readNum(reader)
	bowls := make([][]int, n)
	for i := range n {
		bowls[i] = readNNums(reader, 3)
	}
	return solve(bowls)
}

func solve(bowls [][]int) float64 {
	n := len(bowls)
	if n == 0 {
		return 0
	}

	type bowl struct {
		h float64
		r float64
		R float64
		k float64 // slope: (R-r)/h
	}

	bs := make([]bowl, n)
	for i := 0; i < n; i++ {
		h := float64(bowls[i][0])
		r := float64(bowls[i][1])
		R := float64(bowls[i][2])
		bs[i] = bowl{h: h, r: r, R: R, k: (R - r) / h}
	}

	// y[i] = the y-coordinate of the bottom of bowl i (absolute)
	y := make([]float64, n)
	y[0] = 0

	// requiredOffset computes the minimal d >= 0 such that placing A with its bottom at (B.bottom + d)
	// does not intersect bowl B (same vertical axis).
	//
	// Key observation: radii change linearly with height, so the "A inside B" condition over any overlap
	// interval reduces to checking only the overlap endpoints.
	requiredOffset := func(a, b bowl) float64 {
		const eps = 1e-12
		// If A's bottom doesn't fit into B's top opening, it must be placed above B entirely.
		if a.r >= b.R-eps {
			return b.h
		}

		// Candidate 1: A is entirely within B's height (d + a.h <= b.h).
		best := b.h // worst case: just put on top
		if b.h+eps >= a.h {
			d1 := 0.0
			// Need: r_a <= r_b + k_b*d  => d >= (r_a - r_b)/k_b
			d1 = math.Max(d1, (a.r-b.r)/b.k)
			// Need: R_a <= r_b + k_b*(d+a.h) => d >= (R_a - r_b)/k_b - a.h
			d1 = math.Max(d1, (a.R-b.r)/b.k-a.h)
			if d1 < 0 {
				d1 = 0
			}
			if d1 <= b.h-a.h+1e-10 {
				best = math.Min(best, d1)
			}
		}

		// Candidate 2: A reaches above B's top (d + a.h >= b.h). Overlap is [d, b.h].
		// Need: r_a <= r_b + k_b*d  => d >= (r_a - r_b)/k_b
		// Need at z=b.h: r_a + k_a*(b.h-d) <= R_b => d >= (r_a + k_a*b.h - R_b)/k_a
		d2 := 0.0
		d2 = math.Max(d2, b.h-a.h)
		d2 = math.Max(d2, (a.r-b.r)/b.k)
		d2 = math.Max(d2, (a.r+a.k*b.h-b.R)/a.k)
		if d2 < 0 {
			d2 = 0
		}
		if d2 > b.h {
			d2 = b.h
		}
		best = math.Min(best, d2)
		return best
	}

	for i := 1; i < n; i++ {
		yi := 0.0
		for j := 0; j < i; j++ {
			d := requiredOffset(bs[i], bs[j])
			yi = math.Max(yi, y[j]+d)
		}
		y[i] = yi
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		ans = math.Max(ans, y[i]+bs[i].h)
	}
	return ans
}
