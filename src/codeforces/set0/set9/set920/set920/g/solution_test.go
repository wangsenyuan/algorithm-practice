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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "7 22 1", 9)
}

func TestSample2(t *testing.T) {
	runSample(t, "7 22 2", 13)
}

func TestSample3(t *testing.T) {
	runSample(t, "7 22 3", 15)
}

func TestSample4(t *testing.T) {
	runSample(t, "42 42 42", 187)
}

func TestSample5(t *testing.T) {
	runSample(t, "43 43 43", 87)
}

func TestSample6(t *testing.T) {
	runSample(t, "44 44 44", 139)
}

func TestSample7(t *testing.T) {
	runSample(t, "45 45 45", 128)
}

func TestSample8(t *testing.T) {
	runSample(t, "46 46 46", 141)
}

func TestSample9(t *testing.T) {
	runSample(t, "405280 510510 506878", 3213097)
}
