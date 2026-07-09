package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
5 3
5 2
4 1
5 1
3 2
8
retro
chris
itchy
tuna
crab
rock
cod
ash
`, []string{"Yes", "Yes", "No", "No", "No", "No", "No", "No"})
}

func TestSample2(t *testing.T) {
	runSample(t, `5
5 1
5 2
5 3
5 4
5 5
8
retro
chris
itchy
tuna
crab
rock
cod
ash
`, []string{"Yes", "Yes", "Yes", "No", "No", "No", "No", "No"})
}
