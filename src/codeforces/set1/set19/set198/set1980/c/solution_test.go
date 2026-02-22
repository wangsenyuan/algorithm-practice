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
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 1
1 3 2
4
1 3 1 2`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 2 3 5
2 1 3 5
2
2 3`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
7 6 1 10 10
3 6 1 11 11
3
4 3 11`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
3 1 7 8
2 2 7 10
5
10 3 2 2 1`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5
5 7 1 7 9
4 10 1 2 9
8
1 1 9 8 7 2 10 4`
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `4
1000000000 203 203 203
203 1000000000 203 1000000000
2
203 1000000000`
	expect := false
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `1
1
1
5
1 3 4 5 1`
	expect := true
	runSample(t, s, expect)
}
