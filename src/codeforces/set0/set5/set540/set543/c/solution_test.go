package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 5
abcde
abcde
abcde
abcde
1 1 1 1 1
1 1 1 1 1
1 1 1 1 1
1 1 1 1 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3
abc
aba
adc
ada
10 10 10
10 1 10
10 10 10
10 1 10
`
	expect := 2
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `3 3
abc
ada
ssa
1 1 1
1 1 1
1 1 1
`
	expect := 0
	runSample(t, s, expect)
}
