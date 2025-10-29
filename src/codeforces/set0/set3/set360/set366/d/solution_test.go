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
	s := `4 4
1 2 1 10
2 4 3 5
1 3 1 5
2 4 2 7
	`
	runSample(t, s, 6)
}

func TestSample2(t *testing.T) {
	s := `5 6
1 2 1 10
2 5 11 20
1 4 2 5
1 3 10 11
3 4 12 10000
4 5 6 6
	`
	runSample(t, s, 0)
}

func TestSample3(t *testing.T) {
	s := `5 5
1 5 9403 40347
1 3 13851 29314
4 5 1315 561894
3 5 2748 33090
5 3 10717 32306
	`
	runSample(t, s, 30945)
}
