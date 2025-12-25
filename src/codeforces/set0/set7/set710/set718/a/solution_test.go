package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 1
10.245
`
	expect := "10.25"
	runSample(t, s, expect)
}
func TestSample2(t *testing.T) {
	s := `6 2
10.245
`
	expect := "10.3"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 100
9.2
`
	expect := "9.2"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 10
1.555
`
	expect := "2"
	runSample(t, s, expect)
}
