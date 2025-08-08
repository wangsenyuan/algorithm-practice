package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
i
love
you
<3i<3love<23you<3
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `7
i
am
not
main
in
the
family
<3i<>3am<3the<3<main<3in<3the<3><3family<3
`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
a
b
c
d
<3a<3b<3c<3d
`, false)
}
