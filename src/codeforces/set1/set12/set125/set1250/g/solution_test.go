package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, a, b, d, res := drive(reader)
	if d != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, d)
	}
	if d < 0 {
		return
	}
	var s1, s2 int
	var i int
	for j := range len(a) {
		s1 += a[j]
		s2 += b[j]
		if s1 >= k {
			t.Fatalf("Sample result is invalid, it leads to human loss at %d", j+1)
		}
		if s2 >= k {
			if i < len(res) {
				t.Fatalf("Sample result is invalid, it seems no need to reset all the way to %d", j+1)
			}
			return
		}
		if i < len(res) && j+1 == res[i] {
			if s1 >= s2 {
				s1 -= s2
				s2 = 0
			} else {
				s2 -= s1
				s1 = 0
			}
			i++
		}
	}
	t.Fatalf("Sample result is invalid, human not win after %d rounds", len(a))
}

func TestSample1(t *testing.T) {
	s := `4 17
1 3 5 7
3 5 7 9`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `11 17
5 2 8 2 4 6 1 2 7 2 5
4 6 3 3 5 1 7 4 2 5 3`
	// 2, 4
	// (5, 4) -> (7, 10) -reset-> (0, 3)
	// (8, 6) -> (10, 9) -reset -> (1, 0)
	// (5, 5) -> (11, 6) -> (12, 13) -> (14, 17)) -> win
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 17
6 1 2 7 2 5
1 7 4 2 5 3`
	expect := -1
	runSample(t, s, expect)
}
