package main

import "testing"

func runSample(t *testing.T, s string) {
	var cnt int
	ask := func(x string) []int {
		cnt++
		if cnt > 7 {
			t.Fatalf("Sample asked too much times %d", cnt)
		}
		res := make([]int, 2)
		for i := range 4 {
			if x[i] == s[i] {
				res[0]++
			} else {
				for j := range 4 {
					if x[j] == s[i] {
						res[1]++
						break
					}
				}
			}
		}
		return res
	}

	res := solve(ask)

	if res != s {
		t.Fatalf("Sample expect %s, but got %s", s, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0123"
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := "7923"
	runSample(t, s)
}
