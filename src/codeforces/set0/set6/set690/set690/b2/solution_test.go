package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]Point) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `8
00000000
00000110
00012210
01234200
02444200
01223200
00001100
00000000
5
00000
01210
02420
01210
00000
7
0000000
0122100
0134200
0013200
0002200
0001100
0000000
0
`
	expect := [][]Point{
		{{2, 3}, {2, 4}, {6, 6}, {5, 2}},
		{{2, 2}, {2, 3}, {3, 3}, {3, 2}},
		{{2, 5}, {4, 5}, {4, 2}},
	}
	runSample(t, s, expect)
}
