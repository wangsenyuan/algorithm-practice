package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	cnt, requests, res := drive(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	freq := make([]int, 6)
	for i, v := range res {
		ww := strings.Split(requests[i], ",")

		ok := false
		for _, w := range ww {
			if w == v {
				ok = true
				break
			}
		}

		if !ok {
			t.Fatalf("Sample result %v, not correct, %d-th request %s is not satisfied", res, i, requests[i])
		}

		freq[findIndex(v)]++
	}

	for i, v := range freq {
		if v > cnt[i] {
			t.Fatalf("Sample result %v, not correct, %d-th size %s is not enough", res, i, sizes[i])
		}
	}
}

func TestSample1(t *testing.T) {
	s := `0 1 0 1 1 0
3
XL
S,M
XL,XXL
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1 2 0 1 1
5
S
M
S,M
XXL,XXXL
XL,XXL
`
	expect := false
	runSample(t, s, expect)
}
