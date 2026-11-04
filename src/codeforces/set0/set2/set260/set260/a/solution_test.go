package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	t.Helper()
	var a, b, n int
	fmt.Fscan(strings.NewReader(s), &a, &b, &n)
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !valid(a, b, n, res) {
		t.Fatalf("invalid result %q", res)
	}
}

func valid(a int, b int, n int, res string) bool {
	prefix := fmt.Sprint(a)
	if res == "-1" {
		for d := range 10 {
			if (a*10+d)%b == 0 {
				return false
			}
		}
		return true
	}
	if len(res) != len(prefix)+n || !strings.HasPrefix(res, prefix) {
		return false
	}
	rem := 0
	for _, ch := range res {
		if ch < '0' || ch > '9' {
			return false
		}
		rem = (rem*10 + int(ch-'0')) % b
	}
	return rem == 0
}

func TestSample1(t *testing.T) {
	runSample(t, `5 4 5
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `12 11 1
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `260 150 10
`)
}
