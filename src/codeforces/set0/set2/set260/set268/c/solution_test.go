package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, m, res := drive(reader)
	if len(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, len(res))
	}

	check := func(a, b []int) bool {
		dx := a[0] - b[0]
		dy := a[1] - b[1]
		d := dx*dx + dy*dy
		d1 := int(math.Sqrt(float64(d)))
		return d1*d1 == d
	}

	for i, cur := range res {
		if cur[0]+cur[1] == 0 {
			t.Fatalf("Cant have origin point")
		}
		if cur[0] < 0 || cur[0] > n || cur[1] < 0 || cur[1] > m {
			t.Fatalf("Sample result %v, not correct", res)
		}
		for j := range i {
			if check(res[i], res[j]) {
				t.Errorf("Sample result %v, not correct, as %v %v not valid", res, res[i], res[j])
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 2`
	runSample(t, s, 3)
}

func TestSample2(t *testing.T) {
	s := `4 3`
	runSample(t, s, 4)
}

func TestSample3(t *testing.T) {
	s := `100 30`
	runSample(t, s, 31)
}

func TestSample4(t *testing.T) {
	s := `100 100`
	runSample(t, s, 101)
}
