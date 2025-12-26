package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	m, a, best, changes, res := drive(reader)

	if best != expect[0] || changes != expect[1] {
		t.Fatalf("Sample expect %v, but got %v", expect, []int{best, changes})
	}

	freq := make([]int, m+1)
	var cnt int
	for i, v := range res {
		if a[i] != res[i] {
			cnt++
		}
		if v <= m {
			freq[v]++
		}
	}

	for i := 1; i <= m; i++ {
		if freq[i] < best {
			t.Fatalf("Sample result %v, not correct, %d-th band is not enough", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 2
1 2 3 2
`
	expect := []int{2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 3
1 3 2 2 2 2 1
`
	expect := []int{2, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4
1000000000 100 7 1000000000
`
	expect := []int{1, 4}
	runSample(t, s, expect)
}
