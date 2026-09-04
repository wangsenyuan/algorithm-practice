package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int64) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	if res := drive(reader); res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
1 1
2 2
3 3
4 4
`, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 4
4 1
1 1
4 4
`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `8
7 2
5 7
2 7
1 3
6 7
3 6
7 5
1 6
`, 12)
}

func TestSample4(t *testing.T) {
	runSample(t, `8
6 1
3 6
1 4
1 1
4 2
5 5
3 4
4 1
`, 8)
}

func TestSample5(t *testing.T) {
	runSample(t, `6
5 5
5 4
3 5
1 5
5 3
2 2
`, 4)
}
