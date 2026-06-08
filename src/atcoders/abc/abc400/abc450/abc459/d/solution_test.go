package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if len(res) == len(s) != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}

	if !expect {
		return
	}

	a := []byte(s)
	b := []byte(res)
	slices.Sort(a)
	slices.Sort(b)

	if string(a) != string(b) {
		t.Fatalf("Sample result %s is invalid, it differs from %s", res, s)
	}

	for i := 0; i+1 < len(res); i++ {
		if res[i] == res[i+1] {
			t.Fatalf("Sample result %s is invalid, it has same letters at %d", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "aiiw"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "doodoo"
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "ababacabababacababa"
	expect := true
	runSample(t, s, expect)
}
