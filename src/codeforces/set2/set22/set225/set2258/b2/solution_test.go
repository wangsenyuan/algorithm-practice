package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 4
1 2 3 4 4
`, "6 14 14 14")
}

func TestSample2(t *testing.T) {
	runSample(t, `5 8
1 1 8 8 8
`, "6 12 26 26 26 26 26 26")
}

func TestSample3(t *testing.T) {
	runSample(t, `1 8
6
`, "2 3 6 6 6 6 6 6")
}

func TestSample4(t *testing.T) {
	runSample(t, `7 9
1 7 5 1 7 5 3
`, "7 17 29 29 29 29 29 29 29")
}

func TestSample5(t *testing.T) {
	runSample(t, `4 1
1 1 1 1
`, "4")
}

func TestSample6(t *testing.T) {
	runSample(t, `3 5
3 1 5
`, "3 7 9 9 9")
}
