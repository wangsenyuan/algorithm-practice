package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
89
88
1000
28923845
`
	expect := []string{
		"88",
		"77",
		"99",
		"28923839",
	}
	runSample(t, s, expect)
}
