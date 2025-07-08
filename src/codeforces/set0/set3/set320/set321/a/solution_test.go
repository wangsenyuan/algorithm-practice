package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readString(reader)

	if expect == "Yes" != res {
		t.Errorf("Sample expect %s, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 2
RU
Yes`)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 2
RU
No`)
}

func TestSample3(t *testing.T) {
	runSample(t, `-1 1000000000
LRRLU
Yes`)
}

func TestSample4(t *testing.T) {
	runSample(t, `0 0
D
Yes`)
}

func TestSample5(t *testing.T) {
	runSample(t, `987654321 987654321
UURRDL
Yes`)
}

func TestSample6(t *testing.T) {
	runSample(t, `4 3
UURRDL
Yes`)
}

func TestSample7(t *testing.T) {
	runSample(t, `2 6
RUUUURLDDDL
Yes`)

}
