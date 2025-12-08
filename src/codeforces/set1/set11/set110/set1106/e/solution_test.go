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
	runSample(t, `5 0 2
1 3 4 5
2 5 5 8
`, 13)
}

func TestSample2(t *testing.T) {
	runSample(t, `10 1 6
1 1 2 4
2 2 6 2
3 3 3 3
4 4 4 5
5 5 5 7
6 6 6 9
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `12 2 6
1 5 5 4
4 6 6 2
3 8 8 3
2 9 9 5
6 10 10 7
8 12 12 9
`, 11)
}
