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
	runSample(t, `5 5
..*..
..*..
*****
..*..
..*..
0`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4
****
.*..
.*..
0`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 3
***
*..
*..
*..
0`)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 5
*****
*.*.*
*****
..*.*
..***
0`)
}

func TestSample5(t *testing.T) {
	runSample(t, `1 4
****
0`)
}

func TestSample6(t *testing.T) {
	runSample(t, `5 5
.....
..*..
.***.
..*..
.....
4`)
}

func TestSample7(t *testing.T) {
	runSample(t, `5 3
...
.*.
.*.
***
.*.
1`)
}
