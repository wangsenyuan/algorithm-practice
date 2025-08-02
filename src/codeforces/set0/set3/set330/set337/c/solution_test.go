package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	var expected int
	fmt.Fscan(reader, &expected)

	if res != expected {
		t.Fatalf("Sample expect %d, but got %d", expected, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3 2
3`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5 4 2
6`
	runSample(t, s)
}