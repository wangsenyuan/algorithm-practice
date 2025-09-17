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
	runSample(t, `10 4 6
7 9 11
6 9 13
7 7 7
3 5 6
`, 17)
}

func TestSample2(t *testing.T) {
	runSample(t, `10 7 1
3 4 15
8 9 8
5 6 8
9 10 6
1 4 2
1 4 10
8 10 13
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `10 1 9
5 10 14
`, -1)
}
