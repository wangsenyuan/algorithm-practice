package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)
	if res != expect {
		t.Errorf("Sample %d failed: expect %d, but got %d", 1, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
1 3
1 2
2 3
1 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
1 4
2 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 15
1 5
3 4
2 3
2 4
3 5
4 5
2 5
2 3
1 3
4 5
2 5
4 5
1 2
3 4
1 5
`
	expect := 166
	runSample(t, s, expect)
}
