package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `ab
acxb
cax`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `a
aaaa
aaabbcc`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `a
aaaa
aabbcc`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `ab
baaa
aaaaa`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `a
b
b`
	expect := false
	runSample(t, s, expect)
}