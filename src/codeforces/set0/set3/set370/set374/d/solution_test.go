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
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10 3
1 3 6
-1
1
1
0
0
-1
0
1
-1
1
011`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 1
1
1
-1
Poor stack!`)
}
