package main

import (
	"bufio"
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(bytes.NewBufferString(s))
	res := drive(reader)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
11
`
	expect := "YES"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
010
`
	expect := "NO"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `12
101111111100
`
	expect := "YES"
	runSample(t, s, expect)
}
