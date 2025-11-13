package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, b, best, swap := drive(reader)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}
	if swap[0] > 0 {
		buf := []byte(a)
		buf[swap[0]-1], buf[swap[1]-1] = buf[swap[1]-1], buf[swap[0]-1]
		a = string(buf)
	}

	var sum int
	for i := range a {
		if a[i] != b[i] {
			sum++
		}
	}

	if sum != best {
		t.Fatalf("Sample expect %d, but got %d", best, sum)
	}
}

func TestSample1(t *testing.T) {
	s := `9
pergament
permanent
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
wookie
cookie
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
petr
egor
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
double
bundle
`
	expect := 2
	runSample(t, s, expect)
}
