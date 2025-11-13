package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
IAO'15
IAO'2015
IAO'1
IAO'9
IAO'0`
	expect := []string{
		"2015",
		"12015",
		"1991",
		"1989",
		"1990",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
IAO'9
IAO'99
IAO'999
IAO'9999
`
	expect := []string{
		"1989",
		"1999",
		"2999",
		"9999",
	}
	runSample(t, s, expect)
}
