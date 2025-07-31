package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `AERLAND DERLAND 2:1
DERLAND CERLAND 0:3
CERLAND AERLAND 0:1
AERLAND BERLAND 2:0
DERLAND BERLAND 4:0
`
	expect := "6:0"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `AERLAND DERLAND 2:2
DERLAND CERLAND 2:3
CERLAND AERLAND 1:3
AERLAND BERLAND 2:1
DERLAND BERLAND 4:1
`
	expect := "IMPOSSIBLE"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `BERLAND AQKBSG 7:7
DCVEYFYW AQKBSG 9:3
VTIAYFW AQKBSG 5:9
VTIAYFW BERLAND 3:0
VTIAYFW DCVEYFYW 7:3
`
	expect := "13:12"
	runSample(t, s, expect)
}
