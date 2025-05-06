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
	s := `6
add 4
add 5
add 1
add 2
add 3
sum
`
	expect := []int{3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `14
add 1
add 7
add 2
add 5
sum
add 6
add 8
add 9
add 3
add 4
add 10
sum
del 1
sum
`
	expect := []int{
		5,
		11,
		13,
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1
sum
`
	expect := []int{0}
	runSample(t, s, expect)
}
