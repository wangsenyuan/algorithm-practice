package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	sum, res := drive(reader)

	n := len(res)

	col := make([]int, n)
	for _, row := range res {
		var s1 int
		for j, v := range row {
			col[j] += v
			s1 += v
		}
		if s1 != sum {
			t.Fatalf("Sample result not correct %v", res)
		}
	}
	for _, v := range col {
		if v != sum {
			t.Fatalf("Sample result not correct %v", res)
		}
	}
	var s1, s2 int
	for i := range n {
		s1 += res[i][i]
		s2 += res[i][n-1-i]
	}
	if s1 != sum || s2 != sum {
		t.Fatalf("Sample result not correct %v", res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 3 4 5 6 7 8 9
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3
1 0 -1 0 2 -1 -2 0 1
`
	runSample(t, s)
}

