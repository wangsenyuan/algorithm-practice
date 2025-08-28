package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7
*..P*P*
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
.**PP.*P.*
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `19
**P.*..*..P..*.*P**
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `35
....*..*.*.*.....*.*..P*...*...*...
`
	expect := 36
	runSample(t, s, expect)
}
