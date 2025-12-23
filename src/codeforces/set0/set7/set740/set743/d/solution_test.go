package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool, expectAns int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ok, ans := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}
	if ans != expectAns {
		t.Fatalf("Sample expect %d, but got %d", expectAns, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `8
0 5 -1 4 3 2 6 5
1 2
2 4
2 5
1 3
3 6
6 7
6 8
`
	expect := true
	expectAns := 25
	runSample(t, s, expect, expectAns)
}

func TestSample2(t *testing.T) {
	s := `4
1 -5 1 1
1 2
1 4
2 3
`
	expect := true
	expectAns := 2
	runSample(t, s, expect, expectAns)
}

func TestSample3(t *testing.T) {
	s := `1
-1
`
	expect := false
	expectAns := 2
	runSample(t, s, expect, expectAns)
}

