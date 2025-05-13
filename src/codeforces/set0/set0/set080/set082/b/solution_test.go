package main

import (
	"bufio"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	ans := make([][]int, len(res))

	reader = bufio.NewReader(strings.NewReader(expect))
	for i := range len(res) {
		ans[i] = readVarNums(reader)
	}

	slices.SortFunc(res, func(a, b []int) int {
		return a[0] - b[0]
	})

	slices.SortFunc(ans, func(a, b []int) int {
		return a[0] - b[0]
	})

	for i := range len(res) {
		if !reflect.DeepEqual(res[i], ans[i]) {
			t.Fatalf("Sample expect %v, but got %v", ans, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
3 2 7 4
3 1 7 3
3 5 4 2
3 1 3 5
4 3 1 2 4
2 5 7
`
	expect := `1 7 
2 2 4 
2 1 3 
1 5 
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
5 6 7 8 9 100
4 7 8 9 1
4 7 8 9 2
3 1 6 100
3 2 6 100
2 1 2
`
	expect := `3 7 8 9 
2 6 100 
1 1 
1 2 
`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
2 1 2
2 1 3
2 2 3
`
	expect := `1 1 
1 2 
1 3 
`
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
2 1 2
`
	expect := `1 2
1 1
`
	runSample(t, s, expect)
}
