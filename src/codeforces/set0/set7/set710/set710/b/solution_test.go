package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSample1(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("4\n1 2 3 4\n"))
	if got := drive(reader); got != 2 {
		t.Fatalf("expected 2, got %d", got)
	}
}
