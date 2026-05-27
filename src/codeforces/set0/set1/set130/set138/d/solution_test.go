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
	runSample(t, `2 2
RL
LR
`, "LOSE")
}

func TestSample2(t *testing.T) {
	runSample(t, `2 2
RR
RR
`, "WIN")
}
