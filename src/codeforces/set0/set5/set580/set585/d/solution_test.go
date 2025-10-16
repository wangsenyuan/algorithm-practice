package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if len(expect) == 0 {
		return
	}

	play := func(ans []string) int {
		score := make([]int, 3)

		for i, cur := range a {
			if ans[i] == "MW" {
				score[1] += cur[1]
				score[2] += cur[2]
			} else if ans[i] == "LW" {
				score[0] += cur[0]
				score[2] += cur[2]
			} else {
				score[0] += cur[0]
				score[1] += cur[1]
			}
		}

		return score[0]
	}

	x := play(expect)
	y := play(res)

	if x != y {
		t.Fatalf("Sample expect %v (%d), but got %v (%d)", expect, x, res, y)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 0 0
0 1 0
0 0 1
`
	expect := []string{
		"LM",
		"MW",
		"MW",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
0 8 9
5 9 -2
6 -8 -7
9 4 5
-4 -9 9
-4 5 2
-6 8 -7
`
	expect := []string{
		"LM",
		"MW",
		"LM",
		"LW",
		"MW",
		"LM",
		"LW",
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
1 0 0
1 1 0
`

	runSample(t, s, nil)
}

func TestSample4(t *testing.T) {
	s := `6
1 0 1
1 1 0
0 1 1
0 1 1
1 1 0
1 0 1
`

	expect := []string{
		"LW",
		"LM",
		"MW",
		"MW",
		"LM",
		"LW",
	}

	runSample(t, s, expect)
}
