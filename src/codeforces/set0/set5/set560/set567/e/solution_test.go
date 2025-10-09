package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 7 1 6
1 2 2
1 3 10
2 3 7
2 4 8
3 5 3
4 5 2
5 6 1
`
	expect := []string{
		"YES",
		"CAN 2",
		"CAN 1",
		"CAN 1",
		"CAN 1",
		"CAN 1",
		"YES",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 1 3
1 2 10
2 3 10
1 3 100
`
	expect := []string{
		"YES",
		"YES",
		"CAN 81",
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2 1 2
1 2 1
1 2 2
`
	expect := []string{
		"YES",
		"NO",
	}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 2 1 2
1 2 6
1 2 6
`
	expect := []string{
		"CAN 1",
		"CAN 1",
	}
	runSample(t, s, expect)
}
