package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, ans := range res {
		expect := readNum(reader)
		if expect != ans {
			t.Fatalf("Sample expect %d, but got %d", expect, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 1
1
1
1
1 1 100
100`)
}
func TestSample2(t *testing.T) {
	runSample(t, `1 1
1
1
1
2 1 100
-1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 6
1 8 2 4
3 3 6 1 5 2
3
1 1 1
2 5 10
1 1 6
8
-1
4`)
}

