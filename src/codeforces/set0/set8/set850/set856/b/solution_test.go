package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
aba
baba
aaab`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
aa
a`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
aaaaa
aaa
aaaaaaaaa`
	// aaaaaaaaa
	// aaaaaaa
	// aaaaa  aaaaa
	// aaa aaa aaa
	// a a a
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7
aaaaaa
aaaaaaa
aa
aaaaaaa
aaaaaaa
aaaaa
a`
	expect := 4
	runSample(t, s, expect)
}
