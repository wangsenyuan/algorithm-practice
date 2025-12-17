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
	s := `10 1 4`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 3 10`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1000000000000000000 1 3`
	expect := 666666666666666667
	runSample(t, s, expect)
}
