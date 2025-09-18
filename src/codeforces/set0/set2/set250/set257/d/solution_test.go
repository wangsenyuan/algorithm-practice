package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)
	var sum int
	for i, v := range a {
		if res[i] == '+' {
			sum += v
		} else {
			sum -= v
		}
	}
	if sum < 0 || sum > a[0] {
		t.Fatalf("Sample result %v, not correct, it sums to %d", res, sum)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
1 2 3 5`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
3 3 5`)
}