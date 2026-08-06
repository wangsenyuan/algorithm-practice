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
	runSample(t, `3
1 2 1
`, "YES")
}

func TestSample2(t *testing.T) {
	runSample(t, `2
7 7
`, "YES")
}

func TestSample3(t *testing.T) {
	runSample(t, `6
1 2 3 1 2 3
`, "NO")
}

func TestSample4(t *testing.T) {
	runSample(t, `6
1 1 2 3 2 3
`, "YES")
}

func TestSample5(t *testing.T) {
	runSample(t, `7
1 2 3 1 2 3 4
`, "NO")
}

func TestSample6(t *testing.T) {
	runSample(t, `6
1 2 1 2 1 1
`, "YES")
}

func TestSample7(t *testing.T) {
	runSample(t, `6
1 2 2 3 3 1
`, "NO")
}
