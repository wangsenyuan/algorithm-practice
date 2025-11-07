package main

import "testing"

func runSample(t *testing.T, k int) {
	res := solve(k)

	n := 1 << k
	calc := func(i int, j int) int {
		var sum int
		for l := 0; l < n; l++ {
			u := 1
			if res[i][l] == '*' {
				u = -1
			}
			v := 1
			if res[j][l] == '*' {
				v = -1
			}
			sum += u * v
		}
		return sum
	}

	for i := range res {
		for j := range i {
			w := calc(i, j)
			if w != 0 {
				t.Fatalf("Sample result %v, not correct, %d * %d = %d", res, i, j, w)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4)
}
