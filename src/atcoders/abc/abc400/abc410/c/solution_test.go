package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, x := range res {
		expect := readNum(reader)
		if x != expect {
			t.Fatalf("Sample expect %d, but got %d", expect, x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
2 3
1 2 1000000
3 4
2 2
2 3
3
1
1000000
`)
}
