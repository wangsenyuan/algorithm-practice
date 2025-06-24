package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
3 50000001 50000002
100000012
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 3 99999999 99999994 1000000
303999988
`)
}
