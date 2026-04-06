package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, x, a, res := drive(reader)

	if len(res) != k || res[len(res)-1] > x {
		t.Fatalf("Sample expect %d, but got %d", k, len(res))
	}

	// a is sorted

	play := func(res []int) int {
		maxDist := inf
		for i, j := 0, 0; i < len(res); i++ {
			curPos := res[i]
			for j < len(a) && a[j] <= curPos {
				maxDist = min(maxDist, curPos-a[j])
				j++
			}
			if j < len(a) {
				maxDist = min(maxDist, a[j]-curPos)
			}
		}
		return maxDist
	}

	w := play(expect)
	v := play(res)

	if w != v {
		t.Fatalf("Sample result %v, expect %d, but got %d", res, w, v)
	}
}

func TestSample1(t *testing.T) {
	s := `4 1 4
1 0 2 4`
	expect := []int{3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5 4
0 1 2 3 4`
	expect := []int{0, 1, 2, 3, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1 4
4 0`
	expect := []int{2}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 4 6
2 4 3`
	expect := []int{0, 1, 5, 6}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 2 12
6 12 0`
	expect := []int{3, 9}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `4 3 12
8 12 0 4`
	expect := []int{2, 6, 10}
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `1 1 1000000000
0`
	expect := []int{1000000000}
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `1 1 1000000000
1000000000`
	expect := []int{0}
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `3 4 9
8 7 9`
	expect := []int{0, 1, 2, 3}
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := `3 4 9
2 0 1`
	expect := []int{6, 7, 8, 9}
	runSample(t, s, expect)
}
