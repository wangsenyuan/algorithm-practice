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
	s := `aaba
2
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `aaabbbb
2
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `abracadabra
10
`
	expect := 20
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `ayi
10
`
	expect := 12
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `zumtumtlitf
2
`
	expect := 6
	runSample(t, s, expect)
}
