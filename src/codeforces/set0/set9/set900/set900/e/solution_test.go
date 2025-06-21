package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
bb?a?
1
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `9
ab??ab???
3
2`)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
ab??ab
4
2`)
}

func TestSample4(t *testing.T) {
	runSample(t, `14
?abaa?abb?b?a?
3
3`)
}
