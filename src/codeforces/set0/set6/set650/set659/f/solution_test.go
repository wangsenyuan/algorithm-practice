package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, a, res := drive(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var sum int

	for i, row := range res {
		for j := range row {
			if a[i][j] < res[i][j] {
				t.Fatalf("Sample result %v, not correct, it can't add value to %d %d", res, i, j)
			}
			// a[i][j] >= res[i][j]
			sum += res[i][j]
		}
	}

	if sum != k {
		t.Fatalf("Sample result %v, not correct, sum of result is %d, but expect %d", res, sum, k)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3 35
10 4 9
9 9 7
	`
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := `4 4 50
5 9 1 1
5 1 1 5
5 1 5 5
5 5 7 1
	`
	runSample(t, s, true)
}

func TestSample3(t *testing.T) {
	s := `2 4 12
1 1 3 1
1 6 2 4
	`
	runSample(t, s, false)
}
