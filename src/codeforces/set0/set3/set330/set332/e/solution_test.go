package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	q := drive(reader)
	if q != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, q)
	}
}

func TestSample1(t *testing.T) {
	s := `abacaba
aba
6
`
	expect := "100001"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `abacaba
aba
3
`
	expect := "0"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `     
  
3
`
	expect := "010"
	runSample(t, s, expect)
}
