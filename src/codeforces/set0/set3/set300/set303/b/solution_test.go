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
	runSample(t, "9 9 5 5 2 1\n", []int{1, 3, 9, 7})
}

func TestSample2(t *testing.T) {
	runSample(t, "100 100 52 50 46 56\n", []int{17, 8, 86, 92})
}

func TestReduceByGCD(t *testing.T) {
	runSample(t, "10 10 4 4 4 2\n", []int{0, 1, 10, 6})
}

func TestPinnedToBorder(t *testing.T) {
	runSample(t, "5 5 0 0 2 3\n", []int{0, 0, 2, 3})
}
