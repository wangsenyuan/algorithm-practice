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
	runSample(t, `3
1 2
2 3
`, "1 3 3")
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 2
3 2
4 2
`, "1 3 4 4")
}
