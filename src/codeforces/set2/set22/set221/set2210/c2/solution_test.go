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
	s := `7
1 2 3 4 5 6 7
100 6 12 20 5 2 7`
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
67 67 67
1000000000 1000000000 1000000000`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
8 10 10 12 12 14
1 1 1 1 1 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8
2010 330 550 2 210 385 1001 323
2010 1000 1200 30 500 1000 2000 1000`
	expect := 7
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
1 2 4 3
1 1 1 1`
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `6
2 3 4 5 1 6
10 15 20 25 5 30`
	expect := 6
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `2
1 2
2 100`
	expect := 2
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `5
2860 2860 143 9009 9009
2860 2860 1430 9009 9009`
	expect := 0
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `3
2 60 60
10 60 60`
	expect := 0
	runSample(t, s, expect)
}
