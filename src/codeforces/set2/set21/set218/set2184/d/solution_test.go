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
4 1
`, []int{3})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
4 2
`, []int{2})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
4 3
`, []int{0})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
4 4
`, []int{0})
}

func TestSample5(t *testing.T) {
	runSample(t, `1
4 5
`, []int{0})
}

func TestSample6(t *testing.T) {
	runSample(t, `1
16 5
`, []int{4})
}

func TestSample7(t *testing.T) {
	runSample(t, `1
16 1
`, []int{15})
}
