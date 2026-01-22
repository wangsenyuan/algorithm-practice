package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	fragments, res := drive(reader)
	if len(res) != len(expect) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	for _, cur := range fragments {
		if !strings.Contains(res, cur) {
			t.Fatalf("Sample result %s, not contain %s", res, cur)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
bcd
ab
cdef
`
	expect := "abcdef"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
x
y
z
w
`
	expect := "xyzw"
	runSample(t, s, expect)
}
