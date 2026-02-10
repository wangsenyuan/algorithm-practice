package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)
	if len(res) == 2 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	checkArithmetic := func(arr []int) bool {
		if len(arr) == 0 {
			return false
		}
		if len(arr) <= 2 {
			return true
		}

		d := arr[1] - arr[0]
		for i := 2; i < len(arr); i++ {
			if arr[i]-arr[i-1] != d {
				return false
			}
		}
		return true
	}

	if !checkArithmetic(res[0]) || !checkArithmetic(res[1]) {
		t.Fatalf("Sample result %v, not valid", res)
	}

	for _, v := range a {
		if len(res[0]) > 0 && v == res[0][0] {
			res[0] = res[0][1:]
			continue
		}
		if len(res[1]) > 0 && v == res[1][0] {
			res[1] = res[1][1:]
			continue
		}
		t.Fatalf("Sample result %v, not valid", res)
	}
}

func TestSample1(t *testing.T) {
	s := `6
4 1 2 7 3 10
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 3 -2 -7
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9
-10 -5 0 1 2 3 4 5 6
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
-5 -3 -8 -11 -13 -18 -17 -23 -29 -35
`
	expect := true
	runSample(t, s, expect)
}