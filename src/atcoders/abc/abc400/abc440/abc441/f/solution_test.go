package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 7
2 5
2 5
3 5
3 10
3 20
`
	expect := "BBCBA"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 3
1 1
1 1
1 2
1 2
1 2
1 3
1 3
`
	expect := "CCBBBAA"
	runSample(t, s, expect)
}
