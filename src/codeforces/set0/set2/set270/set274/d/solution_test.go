package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)
	if len(res) > 0 != expect {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	n := len(a)
	buf := make([][]int, n)
	m := len(res)
	for i := range n {
		buf[i] = make([]int, m)
	}
	for i, j := range res {
		j--
		for r := range n {
			buf[r][i] = a[r][j]
		}
	}

	for _, row := range buf {
		prev := -1
		for _, x := range row {
			if x < 0 {
				continue
			}
			if x < prev {
				t.Fatalf("Sample result %v, not correct", res)
			}
			prev = x
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 -1 -1
1 2 1
2 -1 1`
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := `2 3
1 2 2
2 5 4`
	runSample(t, s, true)
}

func TestSample3(t *testing.T) {
	s := `2 3
1 2 3
3 2 1`
	runSample(t, s, false)
}
