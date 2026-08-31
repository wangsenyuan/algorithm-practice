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
	runSample(t, `5 2 3
5 4 3 2 1
`, "YES")
}

func TestSample2(t *testing.T) {
	runSample(t, `6 2 4
2 1 4 3 6 5
`, "NO")
}

func TestSample3(t *testing.T) {
	runSample(t, `4 2 2
1 2 3 4
`, "YES")
}

func TestSample4(t *testing.T) {
	runSample(t, `5 2 3
1 2 3 5 4
`, "YES")
}
