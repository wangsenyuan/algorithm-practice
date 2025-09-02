package main

import (
	"bufio"
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(bytes.NewBufferString(s))
	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 2 1
1 100 1
1 100 100
`
	expect := 99
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `100 100 100
1 1 1
1 1 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `74 89 5
32 76 99
62 95 36
`
	expect := 3529
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `65 6 5
70 78 51
88 55 78
`
	expect := 7027
	runSample(t, s, expect)
}
