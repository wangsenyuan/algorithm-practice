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
	s := `5
1 2 2 2 2
3
2
3
4
`
	expect := []int{5, 5, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 2 3 4 5 6
2
2
3
`
	expect := []int{5, 4}
	runSample(t, s, expect)
}
