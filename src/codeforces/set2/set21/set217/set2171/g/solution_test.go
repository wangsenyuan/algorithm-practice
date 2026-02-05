package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6
1 3 6 4 3 2
3 7 10 4 4 8
`
	expect := []int{17, 827116}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 1
4 3
`
	expect := []int{3, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
2 3 2 5 1
18 13 10 30 7
`
	expect := []int{12, 288}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
5 4 3 6 2
100 125 231 113 107
`
	expect := []int{35, 567812}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
2 2 2 2
2 2 2 2
`
	expect := []int{0, 1}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `4
1 1 1 1
2 2 2 2
`
	expect := []int{1, 1}
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `7
1 1 1 1 1 1 200000
200000 200000 200000 200000 200000 200000 200000
`
	expect := []int{1199994, 0}
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `3
542264 174876 441510
641112 325241 995342
`
	expect := []int{803045, 366998}
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `2
1 1
1 1
`
	expect := []int{0, 1}
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := `2
1 1
1 2
`
	expect := []int{1, 1}
	runSample(t, s, expect)
}