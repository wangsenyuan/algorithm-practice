package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 2
`
	// 1 1, 1 2, 2 2
	// 1 * ？ + 2 * ？ + 1 * ？
	// 感觉任何一个区间被选中的概率 = 1 / (n * n) 
	expect := 1.5
	runSample(t, s, expect)
}
