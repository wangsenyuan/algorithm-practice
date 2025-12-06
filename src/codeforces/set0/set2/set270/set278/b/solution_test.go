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
	s := `5
threehorses
goodsubstrings
secret
primematrix
beautifulyear
`
	expect := "j"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
aa
bdefghijklmn
opqrstuvwxyz
c
`
	expect := "ab"
	runSample(t, s, expect)
}
