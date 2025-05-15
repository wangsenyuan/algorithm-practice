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
		y := readString(reader)
		if y == "Yes" != x {
			t.Fatalf("Sample expect %s, but got %t", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `8 6
1 2 1 3 3 5 2 1
1 3
2 3
2 4
8 8
1 4
5 8
Yes
Yes
No
Yes
No
Yes
`
	runSample(t, s)
}
func TestSample2(t *testing.T) {
	s := `8 1
1 1 2 2 3 3 4 4
1 8
Yes
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `8 3
2 2 1 1 3 3 4 4
1 4
1 5
3 8
Yes
No
Yes
`
	runSample(t, s)
}
