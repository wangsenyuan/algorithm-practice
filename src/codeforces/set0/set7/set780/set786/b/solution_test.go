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
	s := `3 5 1
2 3 2 3 17
2 3 2 2 16
2 2 2 3 3
3 3 1 1 12
1 3 3 17
0 28 12`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 3 1
3 4 1 3 12
2 2 3 4 10
1 2 4 16
0 -1 -1 12`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `10 8 7
1 10 7 366692903
1 4 8 920363557
2 7 5 10 423509459
2 2 5 7 431247033
2 7 3 5 288617239
2 7 3 3 175870925
3 9 3 8 651538651
3 4 2 5 826387883
-1 -1 175870925 288617239 288617239 423509459 0 423509459 423509459 423509459`
	runSample(t, s)
}
