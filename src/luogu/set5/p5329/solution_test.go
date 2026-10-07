package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, want []int) {
	t.Helper()
	got := drive(bufio.NewReader(strings.NewReader(input)))
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "3\naba\n", []int{2, 3, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, "3\naaa\n", []int{1, 2, 3})
}

func TestSample3(t *testing.T) {
	s := `7
aabaaab`
	runSample(t, s, []int{3, 7, 4, 5, 6, 1, 2})
}
