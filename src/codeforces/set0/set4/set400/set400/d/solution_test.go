package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readString(reader)

	if len(res) > 0 != (expect == "Yes") {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}

	if expect != "Yes" {
		return
	}
	k := len(res)
	correct := make([][]int, k)
	for i := range k {
		correct[i] = readNNums(reader, k)
	}

	if !reflect.DeepEqual(res, correct) {
		t.Errorf("Sample expect %v, but got %v", correct, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4 2
1 3
2 3 0
3 4 0
2 4 1
2 1 2
Yes
0 2
2 0
`
	runSample(t, s)
}
func TestSample2(t *testing.T) {
	s := `3 1 2
2 1
1 2 0
Yes
0 -1
-1 0
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 2 2
2 1
1 2 0
2 3 1
Yes
0 1
1 0
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `3 0 2
1 2
No
`
	runSample(t, s)
}
