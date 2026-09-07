package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	t.Skip("solve TODO")
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
-1 -2 -3 -5 -4
`, "0")
}

func TestSample2(t *testing.T) {
	runSample(t, `4
5 7 10 19
`, "0")
}

func TestSample3(t *testing.T) {
	runSample(t, `5
1 -3 2 -1 10
`, "2\n1 3")
}

func TestSample4(t *testing.T) {
	runSample(t, `4
16 -13 -18 -16
`, "0")
}

func TestSample5(t *testing.T) {
	runSample(t, `11
2 -10 -11 3 -10 15 7 18 16 17 -9
`, "6\n6 3 1 5 4 7")
}
