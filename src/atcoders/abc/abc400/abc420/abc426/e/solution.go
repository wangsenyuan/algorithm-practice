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

	for _, v := range drive(reader) {
		fmt.Fprintf(writer, "%.15f\n", v)
	}
}

func drive(reader *bufio.Reader) []float64 {
	var t int
	fmt.Fscan(reader, &t)
	ans := make([]float64, t)
	for i := range t {
		var tsX, tsY, tgX, tgY int
		var asX, asY, agX, agY int
		fmt.Fscan(reader, &tsX, &tsY, &tgX, &tgY)
		fmt.Fscan(reader, &asX, &asY, &agX, &agY)
		ans[i] = solve(tsX, tsY, tgX, tgY, asX, asY, agX, agY)
	}
	return ans
}

func solve(tsX, tsY, tgX, tgY, asX, asY, agX, agY int) float64 {
	getPosAtTime := func(x0 int, y0 int, x1 int, y1 int, t float64) (x float64, y float64) {
		dx := float64(x1 - x0)
		dy := float64(y1 - y0)
		dist := math.Hypot(dx, dy)

		ratio := min(t/dist, 1.0)

		x = float64(x0) + dx*ratio
		y = float64(y0) + dy*ratio
		return
	}

	getDistAtTime := func(t float64) float64 {
		x1, y1 := getPosAtTime(tsX, tsY, tgX, tgY, t)
		x2, y2 := getPosAtTime(asX, asY, agX, agY, t)
		dx := x1 - x2
		dy := y1 - y2
		return math.Sqrt(dx*dx + dy*dy)
	}

	t1 := math.Hypot(float64(tsX-tgX), float64(tsY-tgY))
	t2 := math.Hypot(float64(asX-agX), float64(asY-agY))

	play := func(l float64, r float64) float64 {
		var best float64 = math.MaxFloat64
		for range 100 {
			dist := (r - l) / 3
			m1 := l + dist
			m2 := r - dist
			d1 := getDistAtTime(m1)
			d2 := getDistAtTime(m2)
			if d1 < d2 {
				r = m2
			} else {
				l = m1
			}
			best = min(best, d1, d2)
		}
		return best
	}

	ans1 := play(0, min(t1, t2))
	ans2 := play(min(t1, t2), max(t1, t2))

	return min(ans1, ans2)
}
