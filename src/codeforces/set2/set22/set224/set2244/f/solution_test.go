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
	runSample(t, `2
1
0 1
`, "YES")
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 2 2
0 0 2 1
`, "YES")
}

func TestSample3(t *testing.T) {
	runSample(t, `5
5 5 2 1
0 0 1 2 0
`, "YES")
}

func TestSample4(t *testing.T) {
	runSample(t, `4
1 1 1
0 2 1 3
`, "NO")
}
