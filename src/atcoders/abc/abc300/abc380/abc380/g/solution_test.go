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
	s := `4 2
1 4 2 3`
	expect := 166374061
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 6
7 4 10 5 6 1 8 2 3 9`
	expect := 499122200
	runSample(t, s, expect)
}
