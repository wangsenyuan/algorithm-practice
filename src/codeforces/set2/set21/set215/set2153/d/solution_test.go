package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 1 1 1 1`
	runSample(t, s, 0)
}

func TestSample2(t *testing.T) {
	s := `4
2 100 99 3`
	runSample(t, s, 2)
}

func TestSample3(t *testing.T) {
	s := `5
2 2 5 9 5`
	runSample(t, s, 4)
}

func TestSample4(t *testing.T) {
	s := `6
1 1 1 2 1 2`
	runSample(t, s, 1)
}

func TestSample5(t *testing.T) {
	s := `3
248215438 248215438 785225899`
	runSample(t, s, 537010461)
}

func TestSample6(t *testing.T) {
	s := `3
248215438 701092198 785225899`
	runSample(t, s, 537010461)
}
