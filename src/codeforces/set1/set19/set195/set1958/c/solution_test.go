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
	runSample(t, `1
2 2
`, []int{1})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2 1
`, []int{2})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
10 3
`, []int{10})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
50 36679020707840
`, []int{16})
}
