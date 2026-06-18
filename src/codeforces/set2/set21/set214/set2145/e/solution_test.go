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
	s := `20 25
4
1 22 1 30
1 22 50 30
5
3 1 25
2 23 22
4 10 27
1 21 21
3 20 26
`
	expect := []int{3, 2, 4, 4, 0}
	runSample(t, s, expect)
}

func TestAllValuesWithinLimits(t *testing.T) {
	s := `10 10
1
1
1
1
1 2 2
`
	expect := []int{1}
	runSample(t, s, expect)
}

func TestUpdateExceedsLimits(t *testing.T) {
	s := `10 10
1
1
1
1
1 11 1
`
	expect := []int{0}
	runSample(t, s, expect)
}
