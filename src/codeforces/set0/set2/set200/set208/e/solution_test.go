package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for _, v := range res {
		var expect int
		fmt.Fscan(reader, &expect)
		if v != expect {
			t.Fatalf("Sample expect %d, but got %d", expect, v)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6
0 1 1 0 4 4
7
1 1
1 2
2 1
2 2
4 1
5 1
6 1
0 0 1 0 0 1 1 
`
	runSample(t, s)
}
