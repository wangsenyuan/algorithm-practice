package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, want string) {
	t.Helper()
	got := drive(bufio.NewReader(strings.NewReader(input)))
	if got != want {
		t.Fatalf("expected %s, got %s", want, got)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "6\n1 2 3 4 5 6\n", "YES")
}

func TestSample2(t *testing.T) {
	runSample(t, "5\n3 3 3 3 3\n", "NO")
}

func TestSample3(t *testing.T) {
	runSample(t, "3\n2 2 3\n", "YES")
}

func TestSample4(t *testing.T) {
	runSample(t, "5\n1 1 2 3 7\n", "NO")
}

func TestSample5(t *testing.T) {
	runSample(t, "3\n1 2 2\n", "YES")
}

func TestSample6(t *testing.T) {
	runSample(t, "3\n1 1 2\n", "NO")
}

func TestSample7(t *testing.T) {
	runSample(t, "6\n5 2 10 10 10 2\n", "YES")
}

func TestSample8(t *testing.T) {
	runSample(t, "4\n3 6 9 3\n", "NO")
}
