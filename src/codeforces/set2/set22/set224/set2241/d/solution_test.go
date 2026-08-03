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
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
1 2 3
1 2 3
`, "YES")
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 4 5 2
1 5 4 3
`, "YES")
}

func TestSample3(t *testing.T) {
	runSample(t, `1
9
8
`, "NO")
}

func TestSample4(t *testing.T) {
	runSample(t, `6
6 7 6 7 6 7
7 6 7 6 7 6
`, "YES")
}

func TestSample5(t *testing.T) {
	runSample(t, `9
9 8 7 6 5 4 3 2 1
9 9 8 2 4 4 3 5 3
`, "NO")
}

func TestSample6(t *testing.T) {
	runSample(t, `3
1 1 2
2 1 1
`, "YES")
}

func TestSample7(t *testing.T) {
	runSample(t, `2
1 2
1 1
`, "NO")
}
