package main

import (
	"bufio"
	"slices"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	sum, limit, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	for _, x := range res {
		if x < 1 || x > limit {
			t.Fatalf("Sample result %v, not in range [1, %d]", res, limit)
		}
		sum -= x & -x
	}
	if sum != 0 {
		t.Fatalf("Sample result %v, sum not right", res)
	}

	sort.Ints(res)
	n := len(res)
	res = slices.Compact(res)
	if n != len(res) {
		t.Fatalf("Sample result %v, duplicate elements", res)
	}
}

func TestSample1(t *testing.T) {
	s := "5 5"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "4 3"
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "5 1"
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "54321 12345"
	expect := true
	runSample(t, s, expect)
}
