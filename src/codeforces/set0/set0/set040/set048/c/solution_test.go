package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	unique, res := process(reader)
	expect := readString(reader)
	if unique != (expect == "unique") {
		t.Fatalf("Sample expect %s, but got %v", expect, unique)
	}
	if unique {
		expect_res := readNum(reader)
		if res != expect_res {
			t.Fatalf("Sample expect %d, but got %d", expect_res, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 4
unique
5`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2
1 2
not unique`
	runSample(t, s)
}

func TestSample6(t *testing.T) {
	s := `3
1 3 4
unique
6`
	runSample(t, s)
}
