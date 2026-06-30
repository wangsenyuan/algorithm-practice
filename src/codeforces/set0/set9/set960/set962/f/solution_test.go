package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3
1 2
2 3
3 1
`, []int{1, 2, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `6 7
2 3
3 4
4 2
1 2
1 5
5 6
6 1
`, []int{1, 2, 3, 5, 6, 7})
}

func TestSample3(t *testing.T) {
	runSample(t, `5 6
1 2
2 3
2 4
4 3
2 5
5 3
`, nil)
}

func TestTwoCyclesShareOneVertex(t *testing.T) {
	runSample(t, `5 6
1 2
2 3
3 1
3 4
4 5
5 3
`, []int{1, 2, 3, 4, 5, 6})
}
