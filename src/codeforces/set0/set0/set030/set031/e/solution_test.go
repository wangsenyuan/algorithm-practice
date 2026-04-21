package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	get := func(res string) int {
		var a, b int
		for i, x := range res {
			w := int(s[i] - '0')
			if x == 'H' {
				a = a*10 + w
			} else {
				b = b*10 + w
			}
		}
		return a + b
	}

	w := get(expect)
	v := get(res)

	if w != v {
		t.Fatalf("Sample expect %s(%d), but got %s(%d)", expect, w, res, v)
	}
}

func TestSample1(t *testing.T) {
	s := "1234"
	expect := "HHMM"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "9911"
	expect := "HMHM"
	runSample(t, s, expect)
}
