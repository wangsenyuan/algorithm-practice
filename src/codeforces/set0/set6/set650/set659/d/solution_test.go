package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, in string, expect int) {
	reader := bufio.NewReader(strings.NewReader(in))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	in := `6
0 0
0 1
1 1
1 2
2 2
2 0
0 0
`
	expect := 1
	runSample(t, in, expect)
}

func TestSample2(t *testing.T) {
	in := `16
1 1
1 5
3 5
3 7
2 7
2 9
6 9
6 7
5 7
5 3
4 3
4 4
3 4
3 2
5 2
5 1
1 1
`
	expect := 6
	runSample(t, in, expect)
}
