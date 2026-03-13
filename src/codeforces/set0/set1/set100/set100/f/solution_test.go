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
	s := `2
-1
1
`
	expect := "X^2-1"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1
1
`
	expect := "X^2+2*X+1"
	runSample(t, s, expect)
}
