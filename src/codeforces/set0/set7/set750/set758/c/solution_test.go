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
	expect := readNNums(reader, 3)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 3 8 1 1
3 2 3`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 2 9 4 2
2 1 1`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `5 5 25 4 3
1 1 1`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `100 100 1000000000000000000 100 100
101010101010101 50505050505051 50505050505051`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `1 3 7 1 1
3 2 3`
	runSample(t, s)
}
