package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	shows, channels, res := drive(reader)
	if res[0] != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res[0])
	}

	if expect == 0 {
		return
	}

	u := shows[res[1]-1]
	v := channels[res[2]-1]
	x := max(u[0], v[0])
	y := min(u[1], v[1])
	c := v[2]
	sum := (y - x) * c
	if sum != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, sum)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
7 9
1 4
2 8 2
0 4 1
8 9 3
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1
0 0
1 1 10
`
	expect := 0
	runSample(t, s, expect)
}

// Stress test: many overlapping shows to exercise heap Remove/Swap
func TestStressHeap(t *testing.T) {
	s := `10 5
1 20
2 20
3 20
4 20
5 20
6 20
7 20
8 20
9 20
10 20
5 15 10
10 25 5
1 1 100
`
	reader := bufio.NewReader(strings.NewReader(s))
	shows, channels, res := drive(reader)
	if res[0] < 0 {
		t.Fatalf("unexpected negative efficiency %d", res[0])
	}
	// Verify the returned (video, channel) pair is valid
	if res[0] > 0 {
		u := shows[res[1]-1]
		v := channels[res[2]-1]
		x := max(u[0], v[0])
		y := min(u[1], v[1])
		if y > x {
			sum := (y - x) * v[2]
			if sum != res[0] {
				t.Fatalf("efficiency mismatch: expect %d, got %d", res[0], sum)
			}
		}
	}
}
