package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
1 3
0
0 1 -1
`, [][]int{{2}})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2 4
0 100
-100 -50 0 50
`, [][]int{{150, 200}})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
2 4
0 1000
-100 -50 0 50
`, [][]int{{1000, 200}})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
6 6
20 1 27 100 43 42
100 84 1 24 22 77
`, [][]int{{99, 198, 260, 283}})
}

func TestSample5(t *testing.T) {
	runSample(t, `1
8 2
564040265 -509489796 469913620 198872582 -400714529 553177666 131159391 -20796763
-1000000000 1000000000
`, [][]int{{2000000000, 2027422256}})
}

func TestThreeFeasibleSplitsCanHaveOptimumAtLeftEndpoint(t *testing.T) {
	runSample(t, `1
4 4
16 2 -5 17
14 -18 -19 20
`, [][]int{{39, 71}})
}
