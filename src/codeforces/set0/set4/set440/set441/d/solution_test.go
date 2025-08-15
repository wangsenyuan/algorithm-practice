package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 2 3 4 5
2`, `2
1 2 1 3`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
2 1 4 5 3
2`, `1
1 2`)
}
