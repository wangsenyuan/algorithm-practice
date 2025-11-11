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
	s := `10 5
abacdabcda
7 10 0
5 8 1
1 4 0
3 6 0
7 10 1
`
	expect := "cbcaaaabdd"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 1
agjucbvdfk
1 10 1
`
	expect := "abcdfgjkuv"
	runSample(t, s, expect)
}
