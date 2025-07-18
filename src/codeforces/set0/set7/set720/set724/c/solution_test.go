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
			t.Fatalf("Sample expect %s, but got %v", s, res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3 4
1 1
1 2
2 1
2 2
1
-1
-1
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4 6
1 1
2 1
1 2
2 2
1 3
2 3
1
-1
-1
2
5
-1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `7 4 5
1 3
2 2
5 1
5 3
4 3
13
2
9
5
-1`)
}
