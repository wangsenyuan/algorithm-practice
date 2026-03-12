package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	similarIndex, res := drive(reader)

	next := readString(reader)

	if next == "Brand new problem!" {
		if similarIndex > 0 {
			t.Fatalf("expect %s, but got %d %s", next, similarIndex, res)
		}
		return
	}

	if similarIndex == 0 {
		t.Fatalf("expect %s, but got %d %s", next, similarIndex, res)
	}
	expectIndex, _ := strconv.Atoi(next)

	if expectIndex != similarIndex {
		t.Fatalf("expect %d, but got %d %s", expectIndex, similarIndex, res)
	}

	next = readString(reader)
	if next != res {
		t.Fatalf("expect %s, but got %s", res, next)
	}
}

func TestSample1(t *testing.T) {
	s := `4
find the next palindrome
1
10 find the previous palindrome or print better luck next time
1
[:||||||:]
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3
add two numbers
3
1 add
2 two two
3 numbers numbers numbers
Brand new problem!
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `4
these papers are formulas
3
6 what are these formulas and papers
5 papers are driving me crazy
4 crazy into the night
1
[:||||:]
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `3
add two decimals
5
4 please two decimals add
5 decimals want to be added
4 two add decimals add
4 add one two three
7 one plus two plus three equals six
3
[:|||:]
`
	runSample(t, s)
}
