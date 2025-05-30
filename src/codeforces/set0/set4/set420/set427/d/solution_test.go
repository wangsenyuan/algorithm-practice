package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)
	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `apple
pepperoni
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `lover
driver
1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `bidhan
roy
-1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `testsetses
teeptes
3`)
}

func TestSample5(t *testing.T) {
	runSample(t, `abcdabc
bcdbcdadabacdabc
4`)
}
