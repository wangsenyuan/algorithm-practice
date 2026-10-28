package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	t.Helper()
	var n, k int
	fmt.Fscan(bufio.NewReader(strings.NewReader(s)), &n, &k)
	res := drive(bufio.NewReader(strings.NewReader(s)))
	diff := check(n, res)
	if diff != 2*k {
		t.Fatalf("invalid result, got diff %d for n=%d k=%d", diff, n, k)
	}
}

func check(n int, a []int) int {
	if len(a) != 2*n {
		return -1
	}
	seen := make([]bool, 2*n+1)
	for _, v := range a {
		if v < 1 || v > 2*n || seen[v] {
			return -1
		}
		seen[v] = true
	}
	var absSum, signed int
	for i := 0; i < len(a); i += 2 {
		diff := a[i] - a[i+1]
		if diff < 0 {
			absSum -= diff
		} else {
			absSum += diff
		}
		signed += diff
	}
	if signed < 0 {
		signed = -signed
	}
	return absSum - signed
}

func TestSample1(t *testing.T) {
	runSample(t, `1 0
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 1
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 0
`)
}

func TestSample4(t *testing.T) {
	runSample(t, `50000 24000
`)
}

func TestSample5(t *testing.T) {
	runSample(t, `10 3
`)
}
