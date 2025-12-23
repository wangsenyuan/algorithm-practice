package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool, expectAns []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ok, ans := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}
	if len(ans) != len(expectAns) {
		t.Fatalf("Sample expect %d answers, but got %d", len(expectAns), len(ans))
	}
	slices.Sort(ans)
	slices.Sort(expectAns)
	if !slices.Equal(ans, expectAns) {
		t.Fatalf("Sample expect %v, but got %v", expectAns, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `helloworld
ehoolwlroz
`
	expect := true
	expectAns := []string{"e h", "l o", "d z"}
	runSample(t, s, expect, expectAns)
}

func TestSample2(t *testing.T) {
	s := `hastalavistababy
hastalavistababy
`
	expect := true
	runSample(t, s, expect, nil)
}

func TestSample3(t *testing.T) {
	s := `merrychristmas
christmasmerry
`
	expect := false
	runSample(t, s, expect, nil)
}
