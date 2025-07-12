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
	s := `5 4
4 5
2 4
2 5
2 2
1
4
3
3`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2 1
2 1
1`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `6 7
1 4
3 5
2 3
3 3
3 6
2 1
5 1
5
6
7
7
10
11
8`
	runSample(t, s)
}
