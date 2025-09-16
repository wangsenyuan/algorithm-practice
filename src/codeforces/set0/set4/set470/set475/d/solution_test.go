package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

// runSample runs a test case with the given input string and expected output
func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
2 6 3
5
1
2
3
4
6
`
	expect := []int{1, 2, 2, 0, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
10 20 3 15 1000 60 16
10
1
2
3
4
5
6
10
20
60
1000
`
	expect := []int{14, 0, 2, 2, 2, 0, 2, 2, 1, 1}
	runSample(t, s, expect)
}
