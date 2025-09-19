package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, a, best, res := drive(reader)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %v, which have %d length 1", expect, res, best)
	}
	var cnt int
	for i := range len(a) {
		if a[i] != res[i] {
			cnt++
		}
	}
	if cnt > k {
		t.Fatalf("Sample result %v, makes too much modifications, %d", res, cnt)
	}
}

func TestSample1(t *testing.T) {
	s := `7 1
1 0 0 1 1 0 1
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 2
1 0 0 1 0 1 0 1 0 1
`
	expect := 5
	runSample(t, s, expect)
}
