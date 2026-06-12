package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintf(writer, "%.6f\n", x)
	}
}

func drive(reader *bufio.Reader) []float64 {
	var n, a, d int
	fmt.Fscan(reader, &n, &a, &d)

	cars := make([][]int, n)
	for i := range n {
		cars[i] = make([]int, 2)
		fmt.Fscan(reader, &cars[i][0], &cars[i][1])
	}

	return solve(a, d, cars)
}

func solve(a int, d int, cars [][]int) []float64 {
	n := len(cars)

	// 如果一直加速,多久可以到达d

	t1 := math.Sqrt(float64(d*2) / float64(a))

	play := func(t int, v int) float64 {
		// 在最高速度是v的情况下,最快多久能到达d
		// v = a*t
		t0 := float64(v) / float64(a)
		// 有可能加速没有到v就到达d
		if t0 >= t1 {
			return t1 + float64(t)
		}
		d1 := float64(a) * t0 * t0 / 2
		d2 := float64(d) - d1
		t2 := d2 / float64(v)
		return t2 + t0 + float64(t)
	}

	var prev float64
	ans := make([]float64, n)
	for i, cur := range cars {
		ans[i] = max(prev, play(cur[0], cur[1]))
		prev = ans[i]
	}

	return ans
}
