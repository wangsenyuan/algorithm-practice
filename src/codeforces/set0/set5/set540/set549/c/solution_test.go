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
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 3
346 118 330 1403 5244 480`
	expect := "Daenerys"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 3
5 20 12 7 14 101`
	expect := "Stannis"
	runSample(t, s, expect)
}
