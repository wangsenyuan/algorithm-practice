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
	if len(res) != 1 || res[0] != expect {
		t.Errorf("Sample expect %q, but got %#v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
3
1 2 3
`, "Bob")
}

func TestSample2(t *testing.T) {
	runSample(t, `1
1
1
`, "Alice")
}

func TestSample3(t *testing.T) {
	runSample(t, `1
5
10 3 4 7 4
`, "Alice")
}
