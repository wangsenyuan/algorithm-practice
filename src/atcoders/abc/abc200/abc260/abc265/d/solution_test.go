package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10 5 7 5
1 3 2 2 2 3 1 4 3 2
`, "Yes")
}

func TestSample2(t *testing.T) {
	runSample(t, `9 100 101 100
31 41 59 26 53 58 97 93 23
`, "No")
}

func TestSample3(t *testing.T) {
	runSample(t, `7 1 1 1
1 1 1 1 1 1 1
`, "Yes")
}
