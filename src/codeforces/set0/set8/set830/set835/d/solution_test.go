package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abba"
	expect := []int{6, 1, 0, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abacaba"
	expect := []int{12, 4, 1, 0, 0, 0, 0}
	runSample(t, s, expect)
}
