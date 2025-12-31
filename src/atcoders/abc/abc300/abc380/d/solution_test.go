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
	s := `aB
16
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
`
	expect := "a B A b A b a B A b a B a B A b"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `AnUoHrjhgfLMcDIpzxXmEWPwBZvbKqQuiJTtFSlkNGVReOYCdsay
5
1000000000000000000 123456789 1 987654321 999999999999999999
`
	expect := "K a A Z L"
	runSample(t, s, expect)
}
