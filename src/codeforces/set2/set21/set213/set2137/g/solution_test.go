package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 8 10
1 2
1 3
1 4
2 5
3 6
5 7
2 3
3 4
2 1
1 3
1 4
2 1
2 2
2 3
2 4
2 5
2 6
2 7`
	expect := []bool{
		true,
		false,
		true,
		false,
		false,
		true,
		true,
		true,
	}
	runSample(t, s, expect)
}

func TestPaintingAlreadyLosingBlueNodeDoesNotDoubleCount(t *testing.T) {
	s := `4 3 4
1 2
2 3
1 4
1 3
2 1
1 2
2 1`
	expect := []bool{true, true}
	runSample(t, s, expect)
}
