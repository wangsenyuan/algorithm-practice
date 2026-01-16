package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	w, res := drive(reader)

	a := getFreq(w)
	b := getFreq(res)
	c := getFreq(expect)

	get := func(a []int, b []int) int {
		res := len(res)
		for i := range 26 {
			if a[i] != 0 {
				res = min(res, b[i]/a[i])
			}
		}
		return res
	}

	if get(a, b) != get(a, c) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `?aa?
ab
`, "baab")
}

func TestSample2(t *testing.T) {
	runSample(t, `??b?
za
`, "azbz")
}

func TestSample3(t *testing.T) {
	runSample(t, `abcd
abacaba
`, "abcd")
}
