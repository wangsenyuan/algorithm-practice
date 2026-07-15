package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 1 1 8
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 2 2 6
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 7 4 6
`, 1)
}

func TestSample4(t *testing.T) {
	// a=1,b=1 → period "aabb"; single far position
	runSample(t, `1 1 1000000000 1000000000
`, 1)
}

func TestSample5(t *testing.T) {
	// full long range still only {a,b}
	runSample(t, `1 1 1 1000000000
`, 2)
}

func TestSample6(t *testing.T) {
	runSample(t, `3 1 4 10
`, 4)
}
