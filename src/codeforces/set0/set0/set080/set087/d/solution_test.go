package main

import (
	"bufio"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	x, res := process(reader)
	reader = bufio.NewReader(strings.NewReader(expect))
	y, cnt := readTwoNums(reader)
	if x != y || len(res) != cnt {
		t.Fatalf("Sample expect %d %d, but got %d %v", y, cnt, x, res)
	}
	ans := readNNums(reader, cnt)

	sort.Ints(ans)
	if !reflect.DeepEqual(ans, res) {
		t.Fatalf("Sample expect %v, but got %v", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
2 1 5
`
	expect := `2 1
1 
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 2 1
1 3 5
3 4 2
3 5 3
3 6 4
`
	expect := `16 1
2 
`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9
6 4 72697
9 6 72697
1 6 38220
2 6 38220
6 7 72697
6 5 72697
8 6 72697
3 6 38220
`
	expect := `16 5
1 2 5 6 7 
`
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7
1 2 7485
6 7 50574
3 1 50574
3 4 50574
5 6 58286
6 1 58286
`
	expect := `24 1
6
`
	runSample(t, s, expect)
}
