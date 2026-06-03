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
	s := `5 3
apple
lemon
2 4
1 5
5 5
`
	expect := "lpple"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 5
lemwrbogje
omsjbfggme
5 8
4 8
1 3
6 6
1 4
`
	expect := "lemwrfogje"
	runSample(t, s, expect)
}
