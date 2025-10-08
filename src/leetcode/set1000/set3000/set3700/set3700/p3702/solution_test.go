package p3702

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := removeSubstring(s, k)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "(())", 1, "")
}

func TestSample2(t *testing.T) {
	runSample(t, "(()(", 1, "((")
}

func TestSample3(t *testing.T) {
	runSample(t, "((()))()()()", 3, "()()()")
}

func TestSample4(t *testing.T) {
	runSample(t, "((()())(", 2, "((()())(")
}

func TestSample5(t *testing.T) {
	runSample(t, "(((())())(", 2, "((")
}
