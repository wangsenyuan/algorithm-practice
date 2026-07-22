package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect float64) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if math.Abs(res-expect) > 1e-8 {
		t.Fatalf("Sample expect %.10f, but got %.10f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "4 10 13 10 3 1\n", 4.375)
}

func TestSample2(t *testing.T) {
	runSample(t, "1 4 6 2 2 1\n", -1)
}

func TestSample3(t *testing.T) {
	runSample(t, "3 10 15 17 9 2\n", 11.3333333333)
}

func TestTangentToUpperPost(t *testing.T) {
	runSample(t, "1 9 10 3 6 3\n", 2.25)
}
