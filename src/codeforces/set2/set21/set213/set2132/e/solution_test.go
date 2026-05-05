package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	if !slices.Equal(ans, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4 5
10 20 30
1 2 3 4
0 0 0
3 4 7
3 4 4
1 4 4
2 2 4`
	expect := []int{0, 70, 64, 39, 57}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5 2
500000000 300000000 100000000 900000000 700000000
800000000 400000000 1000000000 600000000 200000000
1 4 3
5 2 6`
	expect := []int{2700000000, 4200000000}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4 1
100 100 20 20
100 100 20 20
4 4 5`
	expect := []int{420}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 3 6
2 363 711
286 121 102
1 1 1
3 1 1
1 2 0
1 3 2
0 1 0
3 3 3`
	expect := []int{711,
		711,
		0,
		997,
		0,
		1360,
	}
	runSample(t, s, expect)
}
