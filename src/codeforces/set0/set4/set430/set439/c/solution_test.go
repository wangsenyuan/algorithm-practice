package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res, _, k, p, a := drive(bufio.NewReader(strings.NewReader(s)))

	if len(res) == k != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, len(res) == k)
	}

	if !expect {
		return
	}

	var arr []int
	cnt := make([]int, 2)

	for _, cur := range res {
		arr = append(arr, cur...)
		var sum int
		for _, v := range cur {
			sum += v
		}
		cnt[sum&1]++
	}

	if cnt[0] != p || cnt[1] != k-p {
		t.Fatalf("Sample result %v, not correct", res)
	}
	slices.Sort(arr)
	slices.Sort(a)
	if !slices.Equal(arr, a) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 5 3
2 6 10 5 9`
	expect := true
	runSample(t, s, expect)
}
