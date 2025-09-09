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
	s := `4 3
2 14 3 4
`
	expect := []int{0, 12, 3, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2
0 2 1 255 254
`
	expect := []int{0, 1, 1, 254, 254}
	runSample(t, s, expect)
}
