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
	expect := readNNums(reader, len(res))
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
3 4 1 2
1 2 3 4`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5
5 4 3 1 2
2 1 3 4 5`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `10
10 9 8 7 6 5 4 3 2 1
2 1 4 3 6 5 8 7 10 9 `
	runSample(t, s)
}
