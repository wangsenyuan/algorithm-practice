package main

import (
	"bufio"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, a, res := process(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	var arr []int
	for i := range n {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			if res[i][l] != res[i][r] {
				t.Fatalf("Sample result not palindrome at row %d, left %d, right %d", i, l, r)
			}
			if res[l][i] != res[r][i] {
				t.Fatalf("Sample result not palindrome at column %d, left %d, right %d", i, l, r)
			}
		}
		arr = append(arr, res[i]...)
	}

	slices.Sort(arr)
	slices.Sort(a)

	if !reflect.DeepEqual(arr, a) {
		t.Fatalf("Sample result not equal to input, got %v, expect %v", arr, a)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 8 8 1 2 2 2 2 2 2 2 2 1 8 8 1
`
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := `3
1 1 1 1 1 3 3 3 3
`
	runSample(t, s, true)
}

func TestSample3(t *testing.T) {
	s := `4
1 2 1 9 8 4 3 8 8 3 4 8 9 2 1 1
`
	runSample(t, s, false)
}

func TestSample4(t *testing.T) {
	s := `1
10
`
	runSample(t, s, true)
}

func TestSample5(t *testing.T) {
	s := `7
5 9 5 4 1 9 8 4 5 1 4 10 7 7 8 4 2 4 4 5 4 4 10 3 4 6 8 1 9 9 5 6 8 7 1 8 6 6 7 5 3 1 1 4 7 2 3 3 8
`
	runSample(t, s, true)
}
