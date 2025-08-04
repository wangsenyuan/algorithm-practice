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
	s := `3 1
2 3 3
1 2 2
3 10 10
10`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1 0
1 2 2
0`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1 2
1 2 2
2`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `2 2
1 3 2
2 4 4
4`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `2 5
1 10 5
3 6 5
5`
	runSample(t, s)
}
