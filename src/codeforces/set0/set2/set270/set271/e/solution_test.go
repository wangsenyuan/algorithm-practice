package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	var expect int
	fmt.Fscan(reader, &expect)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 6
2
11
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 6
7
14
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 10
13 7
36
	`)
}
