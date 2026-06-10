package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 1 2
8
1
2
3
4
5
6
7
8
`
	expect := []int{1, 1, 2, 3, 2, 3, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 30
20 26 3 14 4 4 9
10
31
9
21
23
97
99
30
79
57
3
`
	expect := []int{30, 2, 18, 21, 7, 9, 29, 19, 27, 3}
	runSample(t, s, expect)
}

func TestOriginalPrefixQueries(t *testing.T) {
	s := `4 5
5 1 5 2
4
1
2
3
4
`
	expect := []int{5, 1, 5, 2}
	runSample(t, s, expect)
}
