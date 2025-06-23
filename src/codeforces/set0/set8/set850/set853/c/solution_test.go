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
			t.Errorf("Sample expect %d, but got %d", expect, x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 3
1 2
1 1 1 1
1 1 1 2
1 1 2 2
1
1
1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 2
1 3 2 4
4 1 4 4
1 1 2 3
3
5`)
}

func TestSample3(t *testing.T) {
	runSample(t, `7 7
1 2 3 6 7 4 5
3 2 4 6
3 2 5 4
3 6 7 7
1 1 4 4
2 2 7 6
1 2 6 2
1 2 6 7
17
16
11
16
21
11
21`)
}
