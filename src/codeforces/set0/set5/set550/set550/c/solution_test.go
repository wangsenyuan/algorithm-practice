package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	ok, res := solve(s)
	if ok != expect {
		t.Errorf("Sample expect %t, but got %t", expect, ok)
	}
	if ok {
		var num int
		for i, j := 0, 0; i < len(res); i++ {
			for j < len(s) && s[j] != res[i] {
				j++
			}
			if j == len(s) {
				t.Fatalf("Sample result %s, is not a substring of %s", res, s)
			}
			j++
			num = num*10 + int(res[i]-'0')
			num %= 8
		}
		if num != 0 {
			t.Fatalf("Sample result %s, is not a valid number", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "3454"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "111111"
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "7674"
	expect := true
	runSample(t, s, expect)
}
