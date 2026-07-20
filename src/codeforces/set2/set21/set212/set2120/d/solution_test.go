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
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
1 1 5
`, [][]int{{1, 1}})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2 2 2
`, [][]int{{3, 7}})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
90000 80000 70000
`, [][]int{{299929959, 603196135}})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
100000 100000 100000
`, [][]int{{999899938, 270657473}})
}
