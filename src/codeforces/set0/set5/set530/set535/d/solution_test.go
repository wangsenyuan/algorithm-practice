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
	s := `6 2
ioi
1 3`
	runSample(t, s, 26)
}

func TestSample2(t *testing.T) {
	s := `5 2
ioi
1 2`
	runSample(t, s, 0)
}

func TestSample3(t *testing.T) {
	s := `10 5
ab
1 3 4 6 9`
	runSample(t, s, 0)
}

func TestSample4(t *testing.T) {
	s := `10 5
aa
1 2 3 7 9`
	runSample(t, s, 676)
}
