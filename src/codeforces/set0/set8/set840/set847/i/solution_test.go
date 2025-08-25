package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3 100 140
...
A*.
.B.`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3 2 8
B*.
BB*
BBB`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 4 5 4
..*B
..**
D...`, 7)
}
