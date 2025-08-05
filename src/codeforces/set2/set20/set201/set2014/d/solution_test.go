package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	ans := process(reader)

	if !reflect.DeepEqual(ans, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 1 1
1 2`, []int{1, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `4 1 2
1 2
2 4`, []int{2, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `7 2 3
1 2
1 3
6 7`, []int{1, 4})
}

func TestSample4(t *testing.T) {
	runSample(t, `5 1 2
1 2
3 5`, []int{1, 1})
}

func TestSample5(t *testing.T) {
	runSample(t, `9 2 4
7 9
4 8
1 3
2 3`, []int{3, 4})
}
