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
	s := `5 51
50 150 50 148 150`
	expect := 31
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1000000000
42 42 42`
	expect := -2999999937
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 54321
1 8088 45 1 73 1 9198 4991 1 83`
	expect := -162755
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 100
1 1 1`
	expect := 3
	runSample(t, s, expect)
}
