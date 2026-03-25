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
	s := `2 3 2
.**
X..
`
	expect := "RL"
	runSample(t, s, expect)
}
func TestSample2(t *testing.T) {
	s := `5 6 14
..***.
*...X.
..*...
..*.**
....*.
`
	expect := "DLDDLLLRRRUURU"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3 4
***
*X*
***
`
	expect := "IMPOSSIBLE"
	runSample(t, s, expect)
}
