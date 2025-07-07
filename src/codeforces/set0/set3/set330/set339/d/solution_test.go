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
	s := `2 4
1 6 3 5
1 4
3 4
1 2
1 2
1
3
3
3
`
	runSample(t, s)
}
