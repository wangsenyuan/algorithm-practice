package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}
func TestSample1(t *testing.T) {
	s := `ABCBABA
2
BAAB
ABBA
`
	expect := 1
	runSample(t, s, expect)
}
