package main

import "testing"

func runSample(t *testing.T, s string, ops [][]int) {

	play := func(s string) string {
		buf := []byte(s)
		for _, op := range ops {
			l, r := op[0]-1, op[1]-1
			buf[l], buf[r] = buf[r], buf[l]
		}
		return string(buf)
	}

	ask := func(s string) string {
		return play(s)
	}

	res := solve(play(s), ask)

	if res != s {
		t.Errorf("Sample expect %s, but got %s", s, res)
	}
}

func TestSample1(t *testing.T) {
	s := "yzx"
	ops := [][]int{{1, 2}, {2, 3}}
	runSample(t, s, ops)
}
