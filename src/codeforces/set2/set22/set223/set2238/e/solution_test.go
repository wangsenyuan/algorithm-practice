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
	runSample(t, `4
FTFF
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
TNFTT
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
TFTTTN
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `6
TNNFTF
`, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, `7
TNFNTNF
`, 2)
}

func TestSample6(t *testing.T) {
	runSample(t, `6
NNFFNN
`, 2)
}

func TestSample7(t *testing.T) {
	runSample(t, `7
TNTFNTN
`, 2)
}

func TestSample8(t *testing.T) {
	runSample(t, `1
N
`, 0)
}

func TestSample9(t *testing.T) {
	runSample(t, `5
NNNNN
`, 2)
}

func TestSample10(t *testing.T) {
	runSample(t, `10
NNNTTNNNFN
`, 3)
}

func TestSample11(t *testing.T) {
	runSample(t, `5
FFTFF
`, 1)
}
