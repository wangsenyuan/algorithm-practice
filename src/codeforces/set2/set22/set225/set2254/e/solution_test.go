package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
5
`, []int{5})
}

func TestSample2(t *testing.T) {
	runSample(t, `4
-5 2 1 1
`, nil)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
-3 4 2 -1 1 0
`, []int{1, 1, 3, 2, 6, 3})
}

func TestSample4(t *testing.T) {
	runSample(t, `6
-2 -2 4 1 0 1
`, []int{1, 1, 2, 6, 4, 2})
}

func TestSample5(t *testing.T) {
	runSample(t, `7
0 0 -2 3 0 -1 2
`, []int{2, 1, 1, 1, 1, 4, 2})
}

func TestSample6(t *testing.T) {
	runSample(t, `8
-1 -1 -1 -1 5 0 0 1
`, []int{1, 1, 1, 6, 5, 4, 3, 2})
}

func TestSample7(t *testing.T) {
	runSample(t, `5
1000000000 500000000 750000000 100000000 900000000
`, []int{100000000, 600000000, 1350000000, 2250000000, 3250000000})
}

func TestSample8(t *testing.T) {
	runSample(t, `10
1000000000 -1000000000 500000000 -500000000 1 1 -1 -1 2 -2
`, nil)
}
