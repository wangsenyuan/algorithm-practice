package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	res := solve(s)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `????`
	expect := []string{"00", "01", "10", "11"}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1010"
	expect := []string{"10"}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "1?1"
	expect := []string{"01", "11"}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "111?"
	expect := []string{"11"}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "??????????0????????????????0000000000000"
	expect := []string{"00", "10", "11"}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "?111111?00?"
	expect := []string{"10", "11"}
	runSample(t, s, expect)
}
