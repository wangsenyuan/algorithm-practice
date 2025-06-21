package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, cur := range res {
		expect := readNum(reader)
		if cur != expect {
			t.Fatalf("Sample expect %d, but got %d", expect, cur)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
1 2 1
1 3 5
4 1 7
4 5 3
6 4 2
10
+ 3
+ 1
?
+ 6
?
+ 5
?
- 6
- 3
?
5
14
17
10
	`)
}
