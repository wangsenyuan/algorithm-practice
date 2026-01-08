package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, b, s, res := drive(reader)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}
	buf := []byte(s)

	for _, i := range res {
		buf[i-1] = '1'
	}

	var sum int
	last := -1
	for i := range len(buf) {
		if buf[i] == '1' {
			sum += (i - last - 1) / b
			last = i
		}
	}

	if sum >= a {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 1 2 1
00100
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `13 3 2 3
1000000010001
`
	expect := 2
	runSample(t, s, expect)
}
