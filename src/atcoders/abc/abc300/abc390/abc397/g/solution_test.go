package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3 2
1 2
2 3
1 3
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 4 3
1 2
1 3
3 2
2 4
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 2 1
1 2
1 2
`, 0)
}

func TestParallelRoutes(t *testing.T) {
	runSample(t, `4 4 2
1 2
2 4
1 3
3 4
`, 1)
	runSample(t, `4 4 4
1 2
2 4
1 3
3 4
`, 2)
}

func TestBackEdgeCanDecreaseLevel(t *testing.T) {
	runSample(t, `4 4 3
1 2
2 3
3 2
3 4
`, 3)
}
