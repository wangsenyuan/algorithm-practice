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
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 5 5
010
2 4 3
`, "YES")
}

func TestSample2(t *testing.T) {
	runSample(t, `2 4 3
00
3 4
`, "NO")
}

func TestSample3(t *testing.T) {
	runSample(t, `2 4 2
00
3 3
`, "YES")
}

func TestSample4(t *testing.T) {
	runSample(t, `3 4 2
011
1 1 2
`, "NO")
}

func TestSample5(t *testing.T) {
	runSample(t, `4 2 2
1111
2 2 5 2
`, "NO")
}

func TestSample6(t *testing.T) {
	runSample(t, `1 2 6
0
512
`, "NO")
}
