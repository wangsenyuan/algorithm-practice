package main

import (
	"bufio"
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(bytes.NewBufferString(s))
	res := drive(reader)

	if len(res) > 0 != expect {
		t.Errorf("Sample %s, expect %t, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6
credit
debit
money
rich
bank
capitalism
trap
`
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := `4
aaa
aab
abb
bbb
ba
`
	runSample(t, s, true)
}
