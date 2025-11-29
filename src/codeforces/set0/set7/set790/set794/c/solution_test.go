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
	s := `tinkoff
zscoder
`
	expect := "fzfsirk"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `ioi
imo
`
	expect := "ioi"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `abc
aaa
`
	expect := "aab"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `xxxxxx
xxxxxx
`
	expect := "xxxxxx"
	runSample(t, s, expect)
}
