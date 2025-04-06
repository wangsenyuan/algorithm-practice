package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 0
1 1 4
5 1 4`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 0
1 2 3 4
4 3 2 1`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 0
2 1 1 2
1 2 2 1`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8 0
1 2 3 4 5 6 7 8
8 7 6 5 4 3 2 1`
	expect := 8
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 0
81886062 119066288 136702111
81886073 119066292 136702103`
	// (8, 0, 0) (11, 4, 0)
	// 0, 0, 0

	expect := 2
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3 6
1 1 4
5 1 4`

	expect := 0
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `4 1
1 2 3 4
4 3 2 1`

	expect := 2
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `4 1
2 1 1 2
1 2 2 1`

	expect := 2
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `4 2
2 1 1 2
1 2 2 1`

	expect := 1
	runSample(t, s, expect)
}
