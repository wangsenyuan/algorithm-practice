package main

import "testing"

func runSample(t *testing.T, n int, k int, expect string) {
	res := solve(n, k)
	if res == expect {
		return
	}
	if expect == "Impossible" || res == "Impossible" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	if len(res) != 2*n {
		t.Fatalf("Sample result %s, not having the expected length %d", res, 2*n)
	}
	var level int
	var sum int
	for i := 0; i < len(res); i++ {
		if res[i] == '(' {
			sum += level
			level++
		} else {
			level--
		}
		if level < 0 {
			t.Fatalf("result %s, not a valid bracket sequence", res)
		}
	}
	if level != 0 {
		t.Fatalf("result %s, not a valid bracket sequence", res)
	}
	if sum != k {
		t.Fatalf("result %s, not having the expected nesting sum %d, but got %d", res, k, sum)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, "()(())")
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 6, "(((())))")
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 42, "(((((((()()())))))))")
}
