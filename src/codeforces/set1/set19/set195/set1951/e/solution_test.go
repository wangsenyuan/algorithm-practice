package main

import (
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		if checkPalindrome(x) {
			t.Fatalf("Sample got a palindrome %s", x)
		}
	}
	x := buf.String()

	if x != s {
		t.Fatalf("Sample result %s, not correct", x)
	}
}

func TestSample1(t *testing.T) {
	s := "sinktheyacht"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "lllllllll"
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "uwuowouwu"
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "abababa"
	expect := false
	runSample(t, s, expect)
}
