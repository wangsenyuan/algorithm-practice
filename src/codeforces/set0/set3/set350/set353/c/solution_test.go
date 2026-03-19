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
	runSample(t, `2
3 8
10
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
17 0 10 2 1
11010
`, 27)
}

