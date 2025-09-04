package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	k := readNum(reader)
	if k == 0 {
		if len(res) == 0 {
			return
		}
		t.Fatalf("Sample expect %s, but got %v", s, res)
	}
	expect := readNNums(reader, k)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 4
+ 1
+ 2
- 2
- 1
4
1 3 4 5
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 2
+ 1
- 2
1
3
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2 4
+ 1
- 1
+ 2
- 2
0
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `5 6
+ 1
- 1
- 3
+ 3
+ 4
- 4
3
2 3 5
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `2 4
+ 1
- 2
+ 2
- 1
0
`
	runSample(t, s)
}

func TestSample6(t *testing.T) {
	s := `1 1
+ 1
1
1
`
	runSample(t, s)
}

func TestSample7(t *testing.T) {
	s := `2 1
- 2
2
1 2
`
	runSample(t, s)
}

func TestSample8(t *testing.T) {
	s := `5 5
+ 5
+ 2
+ 3
+ 4
+ 1
1
5
`
	runSample(t, s)
}

func TestSample9(t *testing.T) {
	s := `4 4
+ 2
- 1
- 3
- 2
1
4
`
	runSample(t, s)
}

func TestSample10(t *testing.T) {
	s := `3 3
- 2
+ 1
+ 2
1
3
`
	runSample(t, s)
}
