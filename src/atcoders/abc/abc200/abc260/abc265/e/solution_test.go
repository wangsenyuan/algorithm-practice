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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 2
1 1 1 2 1 3
1 2
2 2
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `10 3
-1000000000 -1000000000 1000000000 1000000000 -1000000000 1000000000
-1000000000 -1000000000
1000000000 1000000000
-1000000000 1000000000
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `300 0
0 0 1 0 0 1
`, 292172978)
}
