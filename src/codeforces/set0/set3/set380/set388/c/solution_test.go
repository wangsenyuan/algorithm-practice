package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNNums(reader, 2)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 100
2 1 10
101 10`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1
9 2 8 6 5 9 4 7 1 3
30 15`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3
3 1 3 2
3 5 4 6
2 8 7
18 18`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `3
3 1000 1000 1000
6 1000 1000 1000 1000 1000 1000
5 1000 1000 1000 1000 1000
7000 7000`
	runSample(t, s)
}
