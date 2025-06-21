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
	n := readNum(reader)
	expect := readNNums(reader, n)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6
5 2 1 1 2 2
2
5 4`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4
1000000000 1000000000 1000000000 1000000000
1
1000000002`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `7
4 10 22 11 12 5 6
7
4 10 22 11 12 5 6`
	runSample(t, s)
}
