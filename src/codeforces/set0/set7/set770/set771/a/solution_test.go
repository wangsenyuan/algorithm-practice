package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readString(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 3
1 3
3 4
1 4
YES`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 4
3 1
2 3
3 4
1 2
NO`)
}
func TestSample3(t *testing.T) {
	runSample(t, `10 4
4 3
5 10
8 9
1 2
YES`)
}
