package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans, res := process(reader)
	a := readNNums(reader, 2)
	b := readNNums(reader, a[0])
	if !reflect.DeepEqual(ans, a) || !reflect.DeepEqual(res, b) {
		t.Errorf("Sample expect %v, %v, but got %v, %v", a, b, ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
4 6 9 3 6
1 3
2 
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5
1 3 5 7 9
1 4
1
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `5
2 3 5 7 11
5 0
1 2 3 4 5 
`
	runSample(t, s)
}
