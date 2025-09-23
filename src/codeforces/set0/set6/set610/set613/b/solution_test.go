package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	A, cf, cm, m, a, best, ans := drive(reader)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}
	x, y := 0, A
	var sum int
	for i, v := range ans {
		if v == A {
			x++
		}
		y = min(y, v)
		sum += v - a[i]
	}
	if sum > m {
		t.Fatalf("Sample result %v, not correct, it uses %d more than %d", ans, sum, m)
	}
	res := x*cf + y*cm
	if res != expect {
		t.Fatalf("Sample result %v, not correct, it uses %d more than %d", ans, sum, m)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 5 10 1 5
1 3 1
`, 12)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 5 10 1 339
1 3 1
`, 35)
}

func TestSample3(t *testing.T) {
	runSample(t, `1 100 1 2 30
71
`, 201)
}
