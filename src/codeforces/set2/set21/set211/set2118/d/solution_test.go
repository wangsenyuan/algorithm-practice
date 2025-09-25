package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	if !reflect.DeepEqual(ans, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
1 4
1 0
3
1 2 3
`
	runSample(t, s, []bool{true, false, true})
}

func TestSample2(t *testing.T) {
	s := `9 4
1 2 3 4 5 6 7 8 9
3 2 1 0 1 3 3 1 1
5
2 5 6 7 8
`
	runSample(t, s, []bool{true, true, true, false, false})
}

func TestSample3(t *testing.T) {
	s := `4 2
1 2 3 4
0 0 0 0
4
1 2 3 4
`
	runSample(t, s, []bool{true, true, false, false})
}

func TestSample4(t *testing.T) {
	s := `3 4
1 2 3
3 1 1
3
1 2 3
`
	runSample(t, s, []bool{true, false, true})
}
