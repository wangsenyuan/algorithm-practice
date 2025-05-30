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
	s := `5 2
1 1 1 1 1
1 1 1 1 1
1 5`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6 7
4 3 5 6 4 4
8 6 0 4 3 4
1 5`
	runSample(t, s)
}
