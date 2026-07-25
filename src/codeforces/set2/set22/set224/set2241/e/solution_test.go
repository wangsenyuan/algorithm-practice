package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
5
1 1 1 1 1
1 2
2 3
2 4
4 5
`, []int64{10})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
10
1 2 3 4 5 6 7 8 9 10
1 3
2 6
6 7
5 4
8 3
3 4
4 6
9 1
10 2
`, []int64{48})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
6
12 6 3 18 9 2
3 4
4 5
2 6
6 1
4 2
`, []int64{0})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
8
3 16 9 1 8 16 4 9
2 1
3 1
4 3
3 5
6 3
4 7
8 1
`, []int64{40})
}
