package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, score, res := drive(reader)
	if score != expect {
		t.Fatalf("expect %d, but got %d", expect, score)
	}
	n := len(a)
	m := len(a[0])
	if n+m != len(res) {
		t.Fatalf("expect %d, but got %d", n+m, len(res))
	}
	row := make([]int, n)
	for r, c, i := 0, 0, 0; i < len(res); i++ {
		if res[i] == 'D' {
			row[r] = c
			r++
		} else {
			c++
			if c > m {
				t.Fatalf("Sample result %s, not valid, it exceeds the column limit %d", res, m)
			}
		}
	}
	sum := []int{0, 0}
	for i := range n {
		for j := range m {
			if j < row[i] {
				sum[0] += a[i][j]
			} else {
				sum[1] += a[i][j]
			}
		}
	}
	if sum[0]*sum[1] != score {
		t.Fatalf("Sample result %s, not correct, score %d, but got %d", res, score, sum[0]*sum[1])
	}
}

func TestSample1(t *testing.T) {
	s := `5 5
1 0 1 1 0
0 1 0 1 1
1 0 1 0 0
0 1 0 1 0
0 0 0 0 1
`
	expect := 30
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4
0 0 1 0
0 1 1 1
1 0 0 1
0 1 0 1
0 0 1 0
`
	expect := 20
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 2
1 0
0 1
1 1
`
	expect := 4
	runSample(t, s, expect)
}
