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
		t.Errorf("Sample %s, expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
2 1 1`
	expect := []int{3, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
3 3 3 5 5`
	expect := []int{10, 9}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
9 9 9 9`
	expect := []int{20, 16}
	runSample(t, s, expect)
}
