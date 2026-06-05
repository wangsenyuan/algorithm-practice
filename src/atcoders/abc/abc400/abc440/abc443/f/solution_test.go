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
	s := `21`
	expect := "126"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10`
	expect := "-1"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1089`
	expect := "9999999999999999999999"
	runSample(t, s, expect)
}
