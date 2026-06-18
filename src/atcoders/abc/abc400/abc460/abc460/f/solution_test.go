package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7
1 2
2 3
2 4
6 7
3 5
7 3
9
1
4
2
6
3
1
1
4
6
`
	expect := []int{4, 3, 3, 2, 2, 3, 2, 3, 4}
	runSample(t, s, expect)
}

func TestLCAAcrossDifferentSubtrees(t *testing.T) {
	s := `4
1 2
2 4
1 3
1
1
`
	expect := []int{3}
	runSample(t, s, expect)
}
