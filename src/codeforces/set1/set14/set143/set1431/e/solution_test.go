package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, b, res := drive(reader)

	check := func(res []int) int {
		ans := 1 << 60
		for i, j := range res {
			ans = min(ans, abs(a[i]-b[j-1]))
		}
		return ans
	}
	x := check(expect)
	y := check(res)

	if x != y {
		t.Fatalf("Sample result %v is incorrect, it gives %d, but expect %d", res, y, x)
	}

	slices.Sort(res)
	for i := range res {
		if res[i] != i+1 {
			t.Fatalf("Sample result %v is invalid, it is not a permuation", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2 3 4
1 2 3 4`
	expect := []int{3, 4, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 100
100 101`
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
1 100
50 51`
	expect := []int{2, 1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
1 1 1 1 1
3 3 3 3 3`
	expect := []int{5, 4, 2, 3, 1}
	runSample(t, s, expect)
}
