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
	if !strings.EqualFold(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
1 4 6 1
3 2 3 7
`, "Mai")
}

func TestSample2(t *testing.T) {
	runSample(t, `6
20 11 1 7 7 0
14 8 3 6 17 6
`, "Ajisai")
}

func TestSample3(t *testing.T) {
	runSample(t, `4
2 6 3 6
3 4 7 1
`, "Tie")
}

func TestSample4(t *testing.T) {
	runSample(t, `5
1 4 5 5 3
6 7 1 2 13
`, "Ajisai")
}

func TestSample5(t *testing.T) {
	runSample(t, `6
9 5 9 17 17 6
1 13 6 13 1 15
`, "Mai")
}

func TestSample6(t *testing.T) {
	runSample(t, `5
2 3 8 1 5
3 1 6 14 7
`, "Tie")
}
