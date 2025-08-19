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
1 3 1
2 4 2
3 4 4
`
	expect := []int{1, 2, 3, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 2
1 5 1
3 7 2
`
	expect := []int{1, 2, 3}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 3
1 1 2
1 1 3
1 1 6
`
	expect := []int{2, 3, 5, 6, 8, 9}
	runSample(t, s, expect)
}
