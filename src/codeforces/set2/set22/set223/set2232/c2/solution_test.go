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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 2 2
EIAIE
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `20 5 5
AEIEEEEIEAAEIEEEEIEA
`, 20)
}

func TestSample3(t *testing.T) {
	runSample(t, `8 2 4
AAAAAIEE
`, 7)
}

func TestSample4(t *testing.T) {
	runSample(t, `8 4 2
AIEAEAAI
`, 7)
}

func TestSample5(t *testing.T) {
	runSample(t, `8 3 3
AIEAEAAI
`, 7)
}

func TestSample6(t *testing.T) {
	runSample(t, `4 2 2
IAEE
`, 4)
}

func TestCannotMoveAnAlreadySeatedAmbivert(t *testing.T) {
	runSample(t, `3 2 2
AEE
`, 2)
}
