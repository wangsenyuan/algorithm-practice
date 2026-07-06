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
	runSample(t, `3 2
0 2
1 2
`, []int{-1, -1, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, `5 5
1 2
1 3
3 4
4 5
0 2
`, []int{3, 3, 3, 3, 2})
}

func TestNeedsTwoSpecialEdges(t *testing.T) {
	runSample(t, `5 4
1 2
0 2
0 4
4 5
`, []int{2, 3, 4, 3, 2})
}

func TestOneSpecialEdgeCanConnectToMiddleTown(t *testing.T) {
	runSample(t, `4 3
1 2
3 4
0 2
`, []int{-1, -1, 3, 2})
}
