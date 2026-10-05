package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 2
1 2 4 5
`, []int{0, 1, 1, 0})
}

func TestSample2(t *testing.T) {
	runSample(t, `4 1
1 2 3 4
`, []int{0, 2, 1, 0})
}

func TestSample3(t *testing.T) {
	runSample(t, `5 7
1 8 9 16 20
`, []int{0, 2, 1, 4, 0})
}

func TestSample4(t *testing.T) {
	runSample(t, `5 1000000000
1 6 7 67 6767
`, []int{0, 0, 0, 0, 0})
}

func TestSample5(t *testing.T) {
	runSample(t, `6 1
1 1 2 2 3 4
`, []int{0, 0, 0, 0, 1, 0})
}

func TestSample6(t *testing.T) {
	runSample(t, `4 1
1 2 3 3
`, []int{0, 1, 0, 0})
}

func TestSample7(t *testing.T) {
	runSample(t, `4 2
1 2 4 6
`, []int{0, 2, 2, 0})
}
