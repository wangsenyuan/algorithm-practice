package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)

	if !reflect.DeepEqual(ans, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `5
4 0 3 1 2`
	expect := []int{3, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 3 4 0`
	expect := []int{3, 4}
	runSample(t, s, expect)
}
