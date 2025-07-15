package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	r, best, ans := process(reader)

	expect := readNum(reader)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	for _, cur := range ans {
		var cnt int
		for i := 0; i < len(cur); i++ {
			if cur[i] == '1' {
				r[i] = max(0, r[i]-1)
				cnt++
			}
		}
		if cnt < 2 || cnt > 5 {
			t.Fatalf("Sample result %v, not correct", ans)
		}
	}
	mn := slices.Min(r)
	mx := slices.Max(r)
	if mn != best || mx != best {
		t.Fatalf("Sample result %v, not correct", ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
4 5 1 7 4
1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
1 2
0`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
1 1 1
1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `100
34 31 31 50 24 48 43 48 22 32 23 22 32 22 32 23 42 20 28 40 32 31 21 52 44 36 29 25 46 20 37 41 36 20 20 46 25 45 26 35 34 25 37 29 38 47 42 25 26 27 48 44 42 45 32 20 25 20 34 50 37 37 20 50 23 27 23 47 39 45 38 23 20 44 48 34 22 49 30 42 24 45 48 28 46 46 42 27 34 23 37 50 39 39 27 44 22 23 34 37
20`)
}
