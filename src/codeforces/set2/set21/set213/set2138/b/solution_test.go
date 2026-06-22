package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 5
1 5 4 3 2
1 2
1 5
3 5
1 4
2 5
`
	expect := []bool{
		true,
		false,
		false,
		false,
		false,
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
3 2 1 4 5
1 1
4 5
1 4
2 5
3 4
`
	expect := []bool{
		true,
		true,
		false,
		true,
		true,
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 1
3 2 4 1
1 4
`

	// 1 + 2 + 1 = 4
	// swap(4, 1) => [3, 2, 1, 4] => swap(3, 1) => [1, 2, 3, 4]

	expect := []bool{false}
	runSample(t, s, expect)
}
