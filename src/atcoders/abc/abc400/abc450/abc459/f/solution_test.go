package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, a []int, expect int) {
	t.Helper()

	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 1, 0}
	runSample(t, a, 3)
}

func TestSample2(t *testing.T) {
	a := []int{4, 6, 3, 5}
	runSample(t, a, 5)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	runSample(t, a, 0)
}

func TestSample4(t *testing.T) {
	a := []int{11, 9, 1, 3, 17, 19, 10, 19, 17, 3}
	runSample(t, a, 78)
}

func TestSampleInput(t *testing.T) {
	s := `4
3
0 1 0
4
4 6 3 5
7
1 2 3 4 5 6 7
10
11 9 1 3 17 19 10 19 17 3
`
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	expect := []int{3, 5, 0, 78}
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestNegativeBlockMergeUsesFloorAndCeil(t *testing.T) {
	a := []int{0, 1, 1}
	runSample(t, a, 2)
}
