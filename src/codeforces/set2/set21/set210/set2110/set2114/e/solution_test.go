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
	s := `5
4 5 2 6 7
1 2
3 2
4 3
5 1
4 5 2 9 7`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6
1000000000 500500500 900900900 9 404 800800800
3 4
5 1
2 5
1 6
6 4
1000000000 1500500096 1701701691 199199209 404 800800800 
`
	runSample(t, s)
}
