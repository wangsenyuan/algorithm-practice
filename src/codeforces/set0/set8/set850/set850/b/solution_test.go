package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 23 17
1 17 17 16
`
	expect := 40
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 6 2
100 49 71 73 66 96 8 60 41 63
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 1000 1
15 17 17 17 17 17 17
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 2 3
1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1 10 1
1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3 10 12
206714 310071 516785
`
	expect := 0
	runSample(t, s, expect)
}

func TestGCD(t *testing.T) {
	nums := []int{206714, 310071, 516785}
	var g int
	for _, num := range nums {
		g = gcd(g, num)
	}
	t.Logf("gcd of %v is %d", nums, g)
}