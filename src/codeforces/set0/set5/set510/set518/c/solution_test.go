package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	var expect int
	fmt.Fscan(reader, &expect)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `8 3 3
1 2 3 4 5 6 7 8
7 8 1
7`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 4 2
3 1 5 2 4
4 4 4 4
8`)
}
