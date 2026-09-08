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
	res, a := drive(reader)

	play := func(arr []int) int {
		slices.Sort(arr)
		var res int
		sign := 1
		j := len(arr) - 1
		for i := len(a) - 1; i >= 0; i-- {
			if j >= 0 && arr[j] == i+1 {
				sign *= -1
				j--
			}
			v := a[i] * sign
			res += v
		}
		return res
	}

	x := play(expect)
	y := play(res)
	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
-1 -2 -3 -5 -4
`, nil)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
5 7 10 19
`, nil)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
1 -3 2 -1 10
`, []int{1, 3})
}

func TestSample4(t *testing.T) {
	runSample(t, `4
16 -13 -18 -16
`, nil)
}

func TestSample5(t *testing.T) {
	runSample(t, `11
2 -10 -11 3 -10 15 7 18 16 17 -9
`, []int{6, 3, 1, 5, 4, 7})
}
