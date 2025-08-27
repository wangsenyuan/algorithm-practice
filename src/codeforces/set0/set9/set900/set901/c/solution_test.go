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
	s := `6 6
1 2
2 3
3 1
4 5
5 6
6 4
3
1 3
4 6
1 6
`
	expect := []int{5, 5, 14}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 9
1 2
2 3
3 1
4 5
5 6
6 7
7 8
8 4
7 2
3
1 8
1 4
3 8
`
	expect := []int{27, 8, 19}
	runSample(t, s, expect)
}
