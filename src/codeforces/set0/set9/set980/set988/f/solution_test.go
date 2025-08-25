package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `10 2 4
3 7
8 10
0 10
3 4
8 1
1 2
	`
	runSample(t, s, 14)
}

func TestSample2(t *testing.T) {
	s := `10 1 1
0 9
0 5
	`
	runSample(t, s, 45)
}

func TestSample3(t *testing.T) {
	s := `10 1 1
0 9
1 5
	`
	runSample(t, s, -1)
}
