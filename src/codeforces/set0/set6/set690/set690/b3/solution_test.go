package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]Point) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `8 19
2 3 2 4 2 5 3 3 3 5 4 3 4 5 4 6 5 2 5 3 5 6 6 2 6 3 6 4 6 5 6 6 6 7 7 6 7 7
5 8
2 2 2 3 2 4 3 2 3 4 4 2 4 3 4 4
0 0
`
	expect := [][]Point{
		{{2, 3}, {2, 4}, {6, 6}, {5, 2}},
		{{2, 2}, {2, 3}, {3, 3}, {3, 2}},
	}
	runSample(t, s, expect)
}
