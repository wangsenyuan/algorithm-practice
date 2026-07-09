package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect [][]int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if len(res) != 1 || !reflect.DeepEqual(res[0], expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
3 5
10101
10100
00101
`, [][]int{
		{6, 6, 6, 9, 9},
		{6, 6, 6, 9, 9},
		{0, 0, 9, 9, 9},
	})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
4 6
011101
010001
100010
101110
`, [][]int{
		{0, 10, 8, 8, 10, 10},
		{0, 10, 8, 8, 10, 10},
		{10, 10, 8, 8, 10, 0},
		{10, 10, 8, 8, 10, 0},
	})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
5 5
11100
10110
11111
01101
00111
`, [][]int{
		{6, 6, 6, 0, 0},
		{6, 6, 4, 4, 0},
		{6, 4, 4, 4, 6},
		{0, 4, 4, 6, 6},
		{0, 0, 6, 6, 6},
	})
}
