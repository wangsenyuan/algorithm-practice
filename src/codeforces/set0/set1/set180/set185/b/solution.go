package main

import "fmt"

func main() {
	var s, a, b, c int
	fmt.Scan(&s, &a, &b, &c)

	res := solve(s, a, b, c)

	fmt.Println(res[0], res[1], res[2])
}

func solve(s int, a int, b int, c int) []float64 {
	if a+b+c == 0 {
		return []float64{0, 0, float64(s)}
	}

	d := float64(a + b + c)

	return []float64{float64(a) * float64(s) / d,
		float64(b) * float64(s) / d,
		float64(c) * float64(s) / d}
}
