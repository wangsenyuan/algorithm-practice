package p3913

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := sortVowels(s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "leetcode"
	expect := "leetcedo"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aeiaaioooa"
	expect := "aaaaoooiie"
	runSample(t, s, expect)
}