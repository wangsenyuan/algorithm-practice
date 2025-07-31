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
	runSample(t, `4
5 5 7 2`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2 3`, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
3 1 10`, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
7`, 1)
}
