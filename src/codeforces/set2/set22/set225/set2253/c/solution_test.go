package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 3 3 3
1 2 3
1 2 3
`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 2 2 2
1 4
2 3
`, 9)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 2 1 1
1
1
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 1 1 5
5
1 2 3 4 5
`, 9)
}

func TestSample5(t *testing.T) {
	runSample(t, `1 1 2 2
1 2
1 2
`, 2)
}

func TestSample6(t *testing.T) {
	runSample(t, `7 2 9 1
1 2 3 4 5 6 7 8 9
9
`, 44)
}

func TestSample7(t *testing.T) {
	runSample(t, `9 9 12 12
1 3 4 6 7 9 10 12 13 15 16 18
2 3 5 6 8 9 11 12 14 15 17 18
`, 170)
}
