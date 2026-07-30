package main

import "testing"

func runSample(t *testing.T, a []int, b []int, p int, expect int) {
	res := solve(a, b, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 5}
	b := []int{6, 1}
	p := 7
	expect := 24
	runSample(t, a, b, p, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1}
	b := []int{1, 1, 1}
	p := 2
	expect := 6
	runSample(t, a, b, p, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2436426, 24979445, 61648772, 23690081, 33933447, 76190629, 62703497}
	b := []int{11047202, 71407775, 28894325, 31963982, 22804784, 50968417, 30302156, 82631932, 61735902, 80895728, 23078537, 7723857}
	p := 25514963
	expect := 2115597124
	runSample(t, a, b, p, expect)
}
