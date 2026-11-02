package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `0 0
0 1
0 2
1 0
1 2
2 0
2 1
2 2
`, "respectable")
}

func TestSample2(t *testing.T) {
	runSample(t, `0 0
1 0
2 0
3 0
4 0
5 0
6 0
7 0
`, "ugly")
}

func TestSample3(t *testing.T) {
	runSample(t, `1 1
1 2
1 3
2 1
2 2
2 3
3 1
3 2
`, "ugly")
}

func TestSample4(t *testing.T) {
	runSample(t, `0 0
2 1
1 0
0 2
2 2
1 0
2 1
0 2
`, "ugly")
}
