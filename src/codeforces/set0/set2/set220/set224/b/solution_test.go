package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, k, res := drive(reader)

	if expect && res[0] == -1 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if !expect && res[0] > 0 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	if !expect {
		return
	}
	// check
	x := slices.Max(a)
	var cnt int
	marked := make([]bool, x+1)
	for i := res[0] - 1; i < res[1]; i++ {
		if !marked[a[i]] {
			cnt++
			marked[a[i]] = true
			if cnt == k && i < res[1]-1 {
				t.Fatalf("Sample result %v, is not correct", res)
			}
		}
	}

}

func TestSample1(t *testing.T) {
	s := `4 2
1 2 2 3
`
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := `8 3
1 1 2 2 3 3 4 5
`
	runSample(t, s, true)
}

func TestSample3(t *testing.T) {
	s := `7 4
4 7 7 4 7 4 7
`
	runSample(t, s, false)
}

func TestSample4(t *testing.T) {
	s := `1 1
5
`
	runSample(t, s, true)
}
