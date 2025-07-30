package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `7 3
abb 2
aaa -3
bba -1
zyz -4
abb 5
aaa 7
xyx 4
`, 12)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1
a 1
a 2
a 3`, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 5
abcde 10000
abcde 10000
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `2 1
t 1
t 2
`, 3)
}

