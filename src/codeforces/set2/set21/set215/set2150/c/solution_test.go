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
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 -1 1
3 1 2
2 3 1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
-2 5 2
3 1 2
2 3 1`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
-1 -2 -3
3 1 2
2 3 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
1000000000 1000000000 1000000000
3 1 2
2 3 1`
	expect := 3000000000
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
5 -15 10 -5
2 4 3 1
1 4 2 3`
	expect := 10
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `4
-5 -5 -5 100
2 3 1 4
4 1 2 3`
	expect := 85
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `4
-1 -100 5 10
1 2 3 4
2 3 4 1`
	expect := 14
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `12
-4 6 10 10 1 -8 6 2 -8 -4 0 -6
11 12 7 3 6 8 1 5 10 2 9 4
7 5 3 6 1 2 8 12 9 4 10 11`
	expect := 24
	runSample(t, s, expect)
}
