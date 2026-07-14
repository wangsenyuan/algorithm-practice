package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 2
`, []int{0, 2, 4, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `6 6
`, []int{1, 31, 90, 65, 15, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `20 5
`, []int{
		0, 0, 0, 331776, 207028224, 204931064, 814022582, 544352515,
		755619435, 401403040, 323173195, 538468102, 309259764, 722947327,
		162115584, 10228144, 423360, 10960, 160, 1,
	})
}
