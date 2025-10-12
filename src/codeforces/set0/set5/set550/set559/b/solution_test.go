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
	s := `aaba
abaa
`
	expect := "YES"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `aabb
abab
`
	expect := "NO"
	runSample(t, s, expect)
}
