package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 0 1000000000
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 0 100
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 1 1000000000
1 4
`, 8)
}

func TestSample4(t *testing.T) {
	runSample(t, `8 8 999999937
1 2
1 3
1 4
1 5
1 6
1 7
1 8
8 7
`, 1)
}
