package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	w, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	if !verify(w, res) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1000
1111
1010
0001`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1111
0111
0010
0111`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
0011
0111
0011
0001`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
1000
0110
0010
1111`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
1000
0110
1010
1111`
	expect := false
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `5
10000
01011
00111
00010
00001`
	expect := false
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `5
10000
11000
10101
10111
00001`
	expect := true
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `5
10000
01101
00100
01110
10001`
	expect := false
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `4
1100
0100
0011
0001`
	expect := false
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := `4
1110
0100
0010
0101`
	expect := true
	runSample(t, s, expect)
}

func TestSample11(t *testing.T) {
	s := `3
100
111
101`
	expect := true
	runSample(t, s, expect)
}
