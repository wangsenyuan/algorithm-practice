package p3734

import "testing"

func runSample(t *testing.T, s string, target string, expect string) {
	res := lexPalindromicPermutation(s, target)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "baba"
	target := "abba"
	expect := "baab"
	runSample(t, s, target, expect)
}

func TestSample2(t *testing.T) {
	s := "baba"
	target := "bbaa"
	expect := ""
	runSample(t, s, target, expect)
}

func TestSample3(t *testing.T) {
	s := "abc"
	target := "abb"
	expect := ""
	runSample(t, s, target, expect)
}

func TestSample4(t *testing.T) {
	s := "aac"
	target := "abb"
	expect := "aca"
	runSample(t, s, target, expect)
}

func TestSample5(t *testing.T) {
	s := "aabb"
	target := "baab"
	expect := ""
	runSample(t, s, target, expect)
}
