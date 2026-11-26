package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
2 3 6 4
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
974153 980051 988027 1
`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `8
12 3 5 45 1 7 63 2
`, 11)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
13
`, 1)
}
