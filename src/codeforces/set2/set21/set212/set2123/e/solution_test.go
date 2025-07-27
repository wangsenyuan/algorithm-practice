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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 0 0 1 2
1 2 4 3 2 1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
3 2 0 4 5 1
1 6 5 4 3 2 1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
1 2 0 1 3 2
1 3 5 4 3 2 1`)
}
func TestSample4(t *testing.T) {
	runSample(t, `4
0 3 4 1
1 3 3 2 1`)
}

func TestSample5(t *testing.T) {
	runSample(t, `5
0 0 0 0 0
1 1 1 1 1 1`)
}
