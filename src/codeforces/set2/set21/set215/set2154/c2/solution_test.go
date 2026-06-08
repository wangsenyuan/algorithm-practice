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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

}

func TestSample1(t *testing.T) {
	s := `2
1 1
1 2`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
4 8
41 67`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
1 1 727 1 1
1 1 1000 1 1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
3 11
1 1`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3
2 7 11
1 6 6`
	expect := 5
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `2
19 1
19 13`
	expect := 32
	runSample(t, s, expect)
}
