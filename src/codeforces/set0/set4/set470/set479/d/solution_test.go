package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, x, y, res := drive(reader)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	if expect == 0 {
		return
	}

	a = append(a, res...)
	sort.Ints(a)
	n := len(a)
	search := func(v int, x int) bool {
		j := sort.SearchInts(a, v+x)
		if j < n && a[j] == v+x {
			return true
		}
		j = sort.SearchInts(a, v-x)
		if j < n && a[j] == v-x {
			return true
		}
		return false
	}

	var flag int
	for _, v := range a {
		if search(v, x) {
			flag |= 1
		}
		if search(v, y) {
			flag |= 2
		}
	}

	if flag != 3 {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 250 185 230
0 185 250
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 250 185 230
0 20 185 250
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 300 185 230
0 300
`
	expect := 2
	runSample(t, s, expect)
}
