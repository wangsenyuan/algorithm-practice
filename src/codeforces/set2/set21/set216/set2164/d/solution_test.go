package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	reader := bufio.NewReader(strings.NewReader(input))
	s, x, ok, res := drive(reader)

	if expect < 0 {
		if ok {
			t.Fatalf("Sample expect %d, but got %v", expect, res)
		}
		return
	}
	if !ok {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	buf := s

	n := len(buf)
	for _, cur := range res {
		for i := 1; i < n; i++ {
			if cur[i] != buf[i-1] && cur[i] != buf[i] {
				t.Fatalf("Sample result %v, not valid", res)
			}
		}
		buf = cur
	}

	if buf != x {
		t.Fatalf("Sample expect %s, but got %s", x, buf)
	}
}

func TestSample1(t *testing.T) {
	s := `4 1
abcd
aabd`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
ab
ab`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 3
abcde
abbcc`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9 1
egcnyeluw
eegccyelw`
	expect := -1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 3
vzvylxxmsy
vvvvvllxxx`
	expect := 3
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `4 6
acba
aaac`
	expect := 2
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `5 7
acabb
aaaca`
	expect := 2
	runSample(t, s, expect)
}
