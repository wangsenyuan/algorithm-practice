package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, b, order, op := drive(reader)
	if len(op) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(op))
	}

	c := slices.Clone(order)
	slices.Sort(c)
	c = slices.Compact(c)
	if len(c) != len(a)+len(b) {
		t.Fatalf("Sample result %v, not correct", order)
	}
	// 在一个里面的顺序，不去检查了
	var state int
	if order[0] <= len(a) {
		state = a[order[0]-1]
	} else {
		state = b[order[0]-len(a)-1]
	}

	var cards []int
	var i int
	for _, sz := range op {
		for i < sz {
			if order[i] <= len(a) {
				cards = append(cards, a[order[i]-1])
			} else {
				cards = append(cards, b[order[i]-len(a)-1])
			}
			i++
		}
		state ^= 1
	}

	if state != 0 {
		t.Fatalf("Sample result %v, not correct", order)
	}

	for i < len(a)+len(b) {
		if order[i] <= len(a) {
			if a[order[i]-1] != state {
				t.Fatalf("Sample result %v, not correct", order)
			}
			cards = append(cards, a[order[i]-1])
		} else {
			if b[order[i]-len(a)-1] != state {
				t.Fatalf("Sample result %v, not correct", order)
			}
			cards = append(cards, b[order[i]-len(a)-1])
		}
		i++
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 0 1
4
1 1 1 1`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 1 1 1 1
5
0 1 0 1 0
`
	expect := 4
	runSample(t, s, expect)
}
