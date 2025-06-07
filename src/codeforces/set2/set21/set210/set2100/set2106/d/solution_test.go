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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `9 5
3 5 2 3 3 5 8 1 2
4 6 2 4 6
6
`)
}
func TestSample2(t *testing.T) {
	runSample(t, `6 3
1 2 6 8 2 1
5 4 3
3
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 3
4 3 5 4 3
7 4 5
7
`)
}

func TestSample4(t *testing.T) {
	runSample(t, `6 3
8 4 2 1 2 5
6 1 4
0
`)
}

func TestSample5(t *testing.T) {
	runSample(t, `5 5
1 2 3 4 5
5 4 3 2 1
-1
`)
}

func TestSample6(t *testing.T) {
	runSample(t, `6 3
1 2 3 4 5 6
9 8 7
-1
`)
}

func TestSample7(t *testing.T) {
	runSample(t, `5 5
7 7 6 7 7
7 7 7 7 7
7
`)
}
