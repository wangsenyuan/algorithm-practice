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
	runSample(t, `1 1 3 4
2026
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2 10 20
31 41 59
`, 11)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 3 3 4
1 2 3 4
`, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, `6 4 9 20
18 27 180 9 45 99
`, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, `7 4 3 5
6 7 14 12 100 78 4
`, 4)
}

func TestSample6(t *testing.T) {
	runSample(t, `9 4 244 353
9982 4435 3998 2443 5399 8244 3539 9824 4353
`, 569)
}

func TestSample7(t *testing.T) {
	runSample(t, `1 1 3 5
8
`, 0)
}
