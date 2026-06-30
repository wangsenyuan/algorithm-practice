package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
4 1
2 3
4 5
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
2 1
3 2
6 1
5 2
4 3
4 1
`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `10
1000000000 1000000000
1000000000 1
-1000000000 1000000000
-1000000000 1
0 1
2 1
1 2
4 1
3 2
4 3
`, 9)
}
