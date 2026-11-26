package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
gennady korotkevich
petr mitrichev
gaoyuan chen
1 2 3
`, false)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
gennady korotkevich
petr mitrichev
gaoyuan chen
3 1 2
`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `2
galileo galilei
nicolaus copernicus
2 1
`, true)
}
