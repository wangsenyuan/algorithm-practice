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
	s := `7 2
bcabcab
`
	expect := "0000011"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `21 2
ababaababaababaababaa
`
	expect := "000110000111111000011"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1
ab
`
	expect := "11"
	runSample(t, s, expect)
}
