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
	runSample(t, `3 1 5 4
1 2 1 3
`, []int{1, 3, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1 5 3
1 3 1
`, []int{2, 3, 2})
}

func TestSample3(t *testing.T) {
	runSample(t, `3 2 5 3
1 3 1
`, []int{1, 2, 2})
}

func TestSample4(t *testing.T) {
	runSample(t, `1 1 1 1
1
`, []int{1})
}

func TestSample5(t *testing.T) {
	runSample(t, `2 1 1 1
2
`, []int{3, 1})
}

func TestSample6(t *testing.T) {
	runSample(t, `2 2 1 1
1
`, []int{1, 3})
}

func TestSample7(t *testing.T) {
	runSample(t, `2 1 2 1
2`, []int{3, 1})
}
