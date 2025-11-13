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
	runSample(t, `3 1 1 2 0 0
3
1 1
2 1
2 3
	`, 11.084259940083)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 0 4 2 2 0
5
5 2
3 0
5 5
3 5
3 3
	`, 33.121375178000)
}
