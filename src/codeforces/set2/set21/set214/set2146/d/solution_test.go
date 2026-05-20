package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	l, r, sum, res := drive(reader)

	if len(res) != r-l+1 {
		t.Fatalf("Sample result %v, expect %v", res, expect)
	}

	calc := func(arr []int) int {
		var sum int
		for i, v := range arr {
			sum += v | (l + i)
		}

		slices.Sort(arr)
		for i, v := range arr {
			if v != l+i {
				t.Fatalf("arr[%d] = %d, expect %d", i, v, l+i)
			}
		}

		return sum
	}

	got := calc(res)
	if sum != got {
		t.Fatalf("Sample sum %d, but permutation gives %d", sum, got)
	}
	if got != calc(expect) {
		t.Fatalf("Sample result %v, expect %v", res, expect)
	}
}

func TestSample1(t *testing.T) {
	s := `0 3`
	expect := []int{3, 2, 1, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 10`
	expect := []int{10, 8, 7, 6, 9}
	runSample(t, s, expect)
}
