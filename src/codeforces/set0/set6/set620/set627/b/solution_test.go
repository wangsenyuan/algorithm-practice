package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2 2 1 8
1 1 2
1 5 3
1 2 1
2 2
1 4 2
1 3 2
2 1
2 3
`
	expect := []int{3, 6, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4 10 1 6
1 1 5
1 5 5
1 3 2
1 5 2
2 1
2 2
`
	expect := []int{7, 1}
	runSample(t, s, expect)
}
