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
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	check := func(x string, y string) bool {
		if len(x) != len(y) {
			return false
		}
		for i := range x {
			if x[i] != y[i] && x[i] != '?' || y[i] == '?' {
				return false
			}
		}
		return true
	}

	cmp := func(x string, y string) int {
		if len(x) < len(y) {
			return -1
		}
		if len(x) > len(y) {
			return 1
		}
		return strings.Compare(x, y)
	}

	for i, cur := range res {
		if !check(a[i], cur) || i > 0 && cmp(res[i-1], cur) >= 0 {
			t.Fatalf("Sample result %v, not correct %s", res, cur)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
?
18
1?`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
??
?`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
12224
12??5
12226
?0000
?00000
`, true)
}
