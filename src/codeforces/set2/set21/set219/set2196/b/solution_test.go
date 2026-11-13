package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int64) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 1 2 100 4
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
2 2 1 1 2 2
`, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, `10
1 1 2 3 4 1 1 7 3 9
`, 10)
}

func TestSample4(t *testing.T) {
	runSample(t, `2
1000000000 1000000000
`, 0)
}
