package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {

	reader := bufio.NewReader(strings.NewReader(s))
	p, x, y, a, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	var sum int
	for _, v := range res {
		sum += v
		if v > p {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
	for _, v := range a {
		sum += v
	}
	if sum > x {
		t.Fatalf("Sample result %v, not correct", res)
	}
	res = append(res, a...)
	slices.Sort(res)
	n := len(res)

	if res[(n+1)/2] < y {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 3 5 18 4
3 5 4
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 3 5 16 4
5 5 5
`, false)
}
