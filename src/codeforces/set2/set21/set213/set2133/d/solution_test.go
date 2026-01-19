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
	runSample(t, `5
3 1 4 1 2`, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 1 1 1`, 1)
}


func TestSample3(t *testing.T) {
	runSample(t, `6
1 2 1 3 5 2`, 7)
}
