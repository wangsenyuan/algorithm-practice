package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(tt *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	l, ramps, best, res := drive(reader)
	if best != expect {
		tt.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	var pos int
	var time int

	for _, i := range res {
		i--
		x, d, t, p := ramps[i][0], ramps[i][1], ramps[i][2], ramps[i][3]
		if pos <= x-p {
			time += x - pos
			time += t
			pos = x + d
			continue
		}
		// pos > x - p
		if x-p < 0 {
			tt.Fatalf("Sample result %v not valid", res)
		}
		time += pos - (x - p)
		time += p
		time += t
		pos = x + d
	}

	time += l - pos

	if time != best {
		tt.Fatalf("Sample expect %d, but got %d", expect, time)
	}
}

func TestSample1(t *testing.T) {
	s := `2 20
5 10 5 5
4 16 1 7
	`
	expect := 15
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 20
9 8 12 6
15 5 1 1
	`
	expect := 16
	runSample(t, s, expect)
}
