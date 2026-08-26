package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	if res := drive(reader); !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1\n0\n", []int64{1})
}

func TestSample2(t *testing.T) {
	runSample(t, "5\n0 4 0 4 14\n", []int64{2, 5, 2, 5, 6})
}

func TestSample3(t *testing.T) {
	runSample(t, "3\n4 0 0\n", []int64{3, 2, 2})
}

func TestSample4(t *testing.T) {
	runSample(t, "3\n0 0 0\n", []int64{1, 1, 1})
}

func TestSample5(t *testing.T) {
	runSample(t, "3\n0 1 1\n", []int64{1, 2, 2})
}

func TestSample6(t *testing.T) {
	runSample(t, "4\n1 1 1 1\n", []int64{-1})
}

func TestSample7(t *testing.T) {
	runSample(t, "7\n0 4 4 4 4 4 9\n", []int64{-1})
}

func TestSample8(t *testing.T) {
	runSample(t, "5\n0 0 0 3 3\n", []int64{1, 1, 1, 2, 2})
}
