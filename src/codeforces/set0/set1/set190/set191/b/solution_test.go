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
	runSample(t, `5 2
8
2 4 5 3 1
2
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 2
8
3 2 4 1 5
5
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 4
1000000000000000
5 4 3 2 1
5
	`)
}

func TestSample4(t *testing.T) {
	runSample(t, `10 6
462
33 98 95 82 91 63 61 51 68 94
1
	`)
}


func TestSample5(t *testing.T) {
	runSample(t, `10 1
55
45 55 61 64 95 95 98 96 65 81
3
	`)
}