package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	v, a, friends, best, res := process(reader)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	var rating int

	var rem int
	for i, u := range a {
		var sum int
		for _, j := range res[i] {
			friend := friends[j-1]
			l, r, f := friend[0]-1, friend[1]-1, friend[2]
			if r < i || i < l {
				t.Fatalf("Sample result %v, not correct, friend %v not stay at day %d", res[i], friend, i)
			}
			sum += f
		}

		if sum > rem+u+v {
			t.Fatalf("player don't have enough food to feed at day %d", i)
		}
		// sum <= rem + u + v
		y := rem + u - v - sum
		rem = min(y, u-v)

		rating += len(res[i])
	}

	if rating != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, rating)
	}
}

func TestSample1(t *testing.T) {
	s := `4 1
3 2 5 4
3
1 3 2
1 4 1
3 4 2
`
	expect := 7
	runSample(t, s, expect)
}
