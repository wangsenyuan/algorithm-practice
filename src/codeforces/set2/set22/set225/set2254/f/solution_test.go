package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect string) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	if res := drive(reader); res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
1 2
1 0
`, "NO")
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 2 4 7
6 7 5 3
`, "YES")
}

func TestSample3(t *testing.T) {
	runSample(t, `4
1 2 4 8
8 4 2 1
`, "YES")
}

func TestSample4(t *testing.T) {
	runSample(t, `4
1 2 3 4
1 2 4 5
`, "NO")
}

func TestSample5(t *testing.T) {
	runSample(t, `4
1 2 0 3
3 3 0 3
`, "NO")
}

func TestSample6(t *testing.T) {
	runSample(t, `6
3 5 6 9 10 12
6 5 3 12 15 9
`, "YES")
}
