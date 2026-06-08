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
	s := `5 2
3 1 4 1 5`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 4
24 3 22 39 4 29`
	expect := 29
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `15 7
220651272 302798780 874479994 657822311 613294668 479624013 241168404 610547619 762548286 256160531 823041612 951553052 226556081 649525901 153805947`
	expect := 1902064780
	runSample(t, s, expect)
}
