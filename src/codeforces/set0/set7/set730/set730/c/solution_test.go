package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	for _, x := range ans {
		expect := readNum(reader)
		if x != expect {
			t.Fatalf("Sample expect %s, but got %v", s, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6 4
4 2
5 4
1 2
3 2
2
4 1 2
3 2 3
6
1 2 6
2 3 7
3 1 2
4 3 8
5 2 5
6 1 10
2
-1
2
2
3
-1`
	runSample(t, s)
}
