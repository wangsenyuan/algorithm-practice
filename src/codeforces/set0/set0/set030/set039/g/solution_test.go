package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "17int f(int n){if (n < 100) return 17;if (n > 99) return 27;}\n", 99)
}

func TestSample2(t *testing.T) {
	runSample(t, "13int f(int n){if (n == 0) return 0;return f(n - 1) + 1;}\n", 13)
}

func TestSample3(t *testing.T) {
	runSample(t, "144int f(int n){if (n == 0) return 0;if (n == 1) return n;return f(n - 1) + f(n - 2);}\n", 24588)
}

func TestLineSeparatedInputAndMaxAnswer(t *testing.T) {
	runSample(t, "1\nint f(int n){return 1;}\n", 32767)
}
