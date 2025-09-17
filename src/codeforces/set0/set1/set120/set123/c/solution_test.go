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
	s := `1 2 1
1 2
`
	expect := `()`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3 1
1 2 3
4 5 6
`
	expect := `(()
())`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 2 2
3 6
1 4
2 5
`
	expect := `()
)(
()`
	runSample(t, s, expect)
}
