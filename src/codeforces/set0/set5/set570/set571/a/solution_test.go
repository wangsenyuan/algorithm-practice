package main

import (
	"bufio"
	"strings"
	"testing"
)


func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 1 1 2`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2 3 1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 2 1 7`
	expect := 0
	runSample(t, s, expect)
}