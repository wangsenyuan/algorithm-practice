package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	var expect [][]int
	for range 4 {
		cur := make([]int, 2)
		fmt.Fscan(reader, &cur[0], &cur[1])
		expect = append(expect, cur)
	}

	get := func(arr [][]int) float64 {
		var res float64
		for i := 0; i < 3; i++ {
			dx := arr[i][0] - arr[i+1][0]
			dy := arr[i][1] - arr[i+1][1]
			res += math.Sqrt(float64(dx*dx + dy*dy))
		}
		return res
	}
	s1 := get(res)
	s2 := get(expect)
	if s1 != s2 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 1
1 1
0 0
1 0
0 1
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `0 10
0 1
0 10
0 0
0 9

`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `10 10
10 9
0 0
10 10
1 0
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `2 2
0 0
2 2
0 2
2 0
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `555 1
555 1
0 0
555 0
0 1
`
	runSample(t, s)
}
