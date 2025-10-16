package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	cnt, _ := drive(reader)

	if cnt != expect {
		t.Errorf("Sample expect %d, but got %d", expect, cnt)
	}
}
func TestSample1(t *testing.T) {
	s := `4
0 0 1 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
0 1 0 1 0
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
1 0 1 1 1 0 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `14
0 1 0 0 0 1 1 0 1 0 1 0 1 0
`
	expect := 3
	runSample(t, s, expect)
}
