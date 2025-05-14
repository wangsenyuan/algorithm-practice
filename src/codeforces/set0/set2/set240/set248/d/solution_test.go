package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 6
HSHSHS
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `14 100
...HHHSSS...SH
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `23 50
HHSS.......SSHHHHHHHHHH
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `34 45
.HHHSSS.........................HS
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `34 44
.HHHSSS.........................HS
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `34 32
.HHHSSS.........................HS
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `162 108
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH............................................................SSSSSSSSSSSSSS
`
	expect := 88
	runSample(t, s, expect)
}
