package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	x, words, res := drive(reader)

	parts := strings.Split(res, " ")

	mem := make(map[string]bool)
	for _, word := range words {
		mem[word] = true
	}

	var buf strings.Builder

	for _, part := range parts {
		if !mem[part] {
			t.Fatalf("Sample result %v, not valid, %s, not found in words", res, part)
		}
		buf.WriteString(change(part))
	}

	y := buf.String()
	y = strings.TrimSpace(y)
	if y != x {
		t.Fatalf("Sample result %v, not valid, %s, not equal to %s", res, y, x)
	}
}

func TestSample1(t *testing.T) {
	s := `30
ariksihsidlihcdnaehsetahgnisol
10
Kira
hates
is
he
losing
death
childish
L
and
Note
`

	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `12
iherehtolleh
5
HI
Ho
there
HeLLo
hello
`

	runSample(t, s)
}
