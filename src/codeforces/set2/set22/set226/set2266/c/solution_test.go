package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	if res := drive(reader); res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) { runSample(t, "4\n0011\n", 0) }
func TestSample2(t *testing.T) { runSample(t, "4\n1000\n", 3) }
func TestSample3(t *testing.T) { runSample(t, "5\n01000\n", 1) }
func TestSample4(t *testing.T) { runSample(t, "8\n01001101\n", 2) }
func TestSample5(t *testing.T) { runSample(t, "7\n0101010\n", 3) }
func TestSample6(t *testing.T) { runSample(t, "7\n0111101\n", 1) }
