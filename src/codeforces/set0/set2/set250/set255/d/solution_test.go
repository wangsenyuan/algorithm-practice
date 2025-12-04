package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)

	if ans != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `6 4 3 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 3 8 10`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 1 2 10`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1000000000 55 60 715189365`
	expect := 37707
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1000000000 81587964 595232616 623563697`
	expect := 17657
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `1 1 1 1`
	expect := 0
	runSample(t, s, expect)
}
