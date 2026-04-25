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
	s := `5
1 3 2 1 2
1 2
2 3
3 4
3 5`
	expect := []int{18, 10, 5, 0, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
1 2 3 1 3 2 1
1 2
2 3
3 4
4 5
4 6
3 7`
	expect := []int{40, 28, 18, 8, 0, 0, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
5 4 3 2 1
1 2
2 3
3 4
4 5`
	expect := []int{20, 10, 4, 1, 0}
	runSample(t, s, expect)
}
