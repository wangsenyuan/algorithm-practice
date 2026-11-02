package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int64) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 4
3 1 4 2
1 1
1 2
2 4
1 4
`, []int64{3, 5, 11, 17})
}

func TestSample2(t *testing.T) {
	// 2 7 1,
	// 2 + 7 + 1 = 10
	// 2 + 1 + 1 = 4
	runSample(t, `5 3
5 2 7 1 4
2 4
1 3
3 5
`, []int64{14, 20, 15})
}

func TestSample3(t *testing.T) {
	runSample(t, `6 4
0 1000000 2 999999 3 4
1 6
2 5
2 2
4 6
`, []int64{2000031, 2000017, 1000000, 1000015})
}

func TestSample4(t *testing.T) {
	runSample(t, `3 1
2 7 1
1 3
`, []int64{14})
}
