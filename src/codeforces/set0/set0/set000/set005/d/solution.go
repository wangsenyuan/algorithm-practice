package main

import (
	"fmt"
	"math"
)

func main() {
	var A, V, L, D, W int
	fmt.Scanf("%d %d\n", &A, &V)
	fmt.Scanf("%d %d %d\n", &L, &D, &W)
	res := solve(A, V, L, D, W)
	fmt.Printf("%.10f\n", res)
}

const eps = 1e-9

func solve(A int, V int, L int, D int, W int) float64 {
	// 在长度为l的路上，距离起点距离为d的地方，设置一个限速为w的限速
	// 车辆从0开始，加速度为a，最大速度为v，经过限速后，开完全程的时间
	a := float64(A)
	v := float64(V)
	l := float64(L)
	d := float64(D)
	w := float64(W)

	dist := func(speed float64, t float64) float64 {
		return speed*t + a*t*t/2
	}

	travel := func(d float64, speed float64) float64 {
		// a * t * t / 2 + speed * t - d = 0
		delta := speed*speed + 2*a*d
		t := (-speed + math.Sqrt(delta)) / a
		tx := (v - speed) / a
		if t < tx+eps {
			return t
		}
		// 先加速到v，
		return tx + (d-dist(speed, tx))/v
	}

	if v <= w+eps {
		return travel(l, 0)
	}

	tw := w / a
	dw := dist(0, tw)
	if d <= dw+eps {
		return travel(l, 0)
	}

	return tw + 2*travel((d-dw)/2, w) + travel(l-d, w)
}
