package main

import (
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	res := solve(s)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aba,123;1a;0"
	expect := []string{"\"123,0\"", "\"aba,1a\""}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1;;01,a0,"
	expect := []string{"\"1\"", "\",01,a0,\""}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "1"
	expect := []string{"\"1\"", "-"}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := ",;,,;"
	expect := []string{"-", "\",,,,,\""}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "6;2,"
	expect := []string{"\"6,2\"", "\"\""}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := strings.Repeat(",", 100000)
	expect := []string{"-", "\"" + s + "\""}
	runSample(t, s, expect)
}
