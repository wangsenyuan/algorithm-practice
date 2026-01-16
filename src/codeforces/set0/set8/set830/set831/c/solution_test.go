package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 1
-5 5 0 20
10`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 2
-2000 -2000
3998000 4000000
`, 1)
}