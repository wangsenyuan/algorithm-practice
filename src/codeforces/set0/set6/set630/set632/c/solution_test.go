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
	s := `4
abba
abacaba
bcd
er
`
	expect := "abacabaabbabcder"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
x
xx
xxa
xxaa
xxaaa
`
	expect := "xxaaaxxaaxxaxxx"
	runSample(t, s, expect)
}
