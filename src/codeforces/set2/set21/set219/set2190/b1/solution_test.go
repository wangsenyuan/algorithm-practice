package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
()
`, -1)
}

func TestSample2(t *testing.T) {
	runSample(t, `8
(()(()))
`, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
(())()
`, -1)
}
