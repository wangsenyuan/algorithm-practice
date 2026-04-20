package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	x0, z0 := drive(reader)
	fmt.Printf("%.10f %.10f\n", x0, z0)
}

func drive(reader *bufio.Reader) (x0 float64, z0 float64) {
	var a, b, m, vx, vy, vz int
	fmt.Fscan(reader, &a, &b, &m, &vx, &vy, &vz)
	return solve(a, b, m, vx, vy, vz)
}

func solve(a int, b int, m int, vx int, vy int, vz int) (x0 float64, z0 float64) {
	t := float64(m) / float64(-vy)
	dx := float64(vx) * t
	// dx > 0 or dx <= 0

	if math.Abs(dx)*2 <= float64(a) {
		x0 = dx + float64(a)/2
	} else {
		sign := 1
		if dx < 0 {
			dx = -dx
			sign = -1
		}
		dx -= float64(a) / 2
		cnt := int(dx / float64(a))
		dx -= float64(cnt) * float64(a)
		if cnt%2 == 0 {
			x0 = float64(a) - dx
		} else {
			x0 = dx
		}
		if sign == -1 {
			x0 = float64(a) - x0
		}
	}

	dz := float64(vz) * t
	cnt := int(dz / float64(b))
	dz -= float64(cnt) * float64(b)
	if cnt%2 == 0 {
		z0 = dz
	} else {
		z0 = float64(b) - dz
	}
	return
}
