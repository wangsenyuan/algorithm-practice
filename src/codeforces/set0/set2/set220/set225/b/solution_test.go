package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	x, _, res := drive(reader)
	var sum int
	for _, v := range res {
		sum += v
	}
	if sum != x {
		t.Errorf("Sample expect %d, but got %d", x, sum)
	}
}

func TestSample1(t *testing.T) {
	s := `21 5`
	runSample(t, s)
}
