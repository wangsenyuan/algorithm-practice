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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 3
0.0.0.1
0.1.1.2
0.0.2.1
0.1.1.0
0.0.2.3
`, "255.255.254.0")
}

func TestSample2(t *testing.T) {
	runSample(t, `5 2
0.0.0.1
0.1.1.2
0.0.2.1
0.1.1.0
0.0.2.3
`, "255.255.0.0")
}

func TestSample3(t *testing.T) {
	runSample(t, `2 1
255.0.0.1
0.0.0.2
`, "-1")
}
