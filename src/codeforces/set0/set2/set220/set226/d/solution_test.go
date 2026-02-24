package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)
	n := len(a)
	m := len(a[0])
	row := make([]int, n)
	col := make([]int, m)
	for _, i := range res[0] {
		row[i-1] = 1
	}
	for _, i := range res[1] {
		col[i-1] = 1
	}
	s1 := make([]int, n)
	s2 := make([]int, m)
	for i := range n {
		for j := range m {
			if row[i]^col[j] == 1 {
				s1[i] -= a[i][j]
				s2[j] -= a[i][j]
			} else {
				s1[i] += a[i][j]
				s2[j] += a[i][j]
			}
		}
	}
	for i := range n {
		if s1[i] < 0 {
			t.Fatalf("Sample result not correct %v", res)
		}
	}
	for i := range m {
		if s2[i] < 0 {
			t.Fatalf("Sample result not correct %v", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 1
-1
-1
-1
-1
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2 4
-1 -1 -1 2
1 1 1 1
`
	runSample(t, s)
}
