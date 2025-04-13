package p3518

import (
	"testing"
)

func runSample(t *testing.T, s string, k int, expect string) {
	res := smallestPalindrome(s, k)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abba"
	k := 2
	expect := "baab"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "aa"
	k := 2
	expect := ""
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "bacab"
	k := 1
	expect := "abcba"
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "xxnfnxx"
	k := 3
	expect := "xxnfnxx"
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "wewcehhecwew"
	k := 47
	expect := "eehwcwwcwhee"
	runSample(t, s, k, expect)
}
