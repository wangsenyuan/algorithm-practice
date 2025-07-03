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
	runSample(t, `3 1
2 2
0
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 0
1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 3
3 1
3 2
3 3
1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 1
3 2
4`)
}

func TestSample5(t *testing.T) {
	runSample(t, `1000 0
1996`)
}

func TestSample6(t *testing.T) {
	runSample(t, `6 5
2 1
6 4
2 2
4 3
4 1
3`)
}

func TestSample7(t *testing.T) {
	runSample(t, `6 5
2 6
5 2
4 3
6 6
2 5
2`)
}
