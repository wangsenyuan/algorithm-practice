package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	x := strings.Join(res, "\n")

	x = strings.TrimSpace(x)

	expect = strings.TrimSpace(expect)

	if x != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, x)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
aaaaa
acacaca
aabaa
ccacacc
caaac
`
	expect := `YES
NO
NO`
	runSample(t, s, expect)
}
