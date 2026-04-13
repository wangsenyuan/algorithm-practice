package p3899

import (
	"math"
	"slices"
)

func internalAngles(sides []int) []float64 {
	sorted := append([]int(nil), sides...)
	slices.Sort(sorted)

	for i := 0; i+2 < len(sorted); i++ {
		a, b, c := float64(sorted[i]), float64(sorted[i+1]), float64(sorted[i+2])
		if a+b <= c {
			continue
		}
		return anglesOppositeSides(a, b, c)
	}
	return nil
}

// anglesOppositeSides returns interior angles (degrees) opposite sides a, b, c respectively.
func anglesOppositeSides(a, b, c float64) []float64 {
	return []float64{
		acosDeg((b*b + c*c - a*a) / (2 * b * c)),
		acosDeg((a*a + c*c - b*b) / (2 * a * c)),
		acosDeg((a*a + b*b - c*c) / (2 * a * b)),
	}
}

func acosDeg(x float64) float64 {
	if x > 1 {
		x = 1
	}
	if x < -1 {
		x = -1
	}
	return math.Acos(x) * 180 / math.Pi
}
