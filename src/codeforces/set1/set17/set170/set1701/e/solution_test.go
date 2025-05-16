package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	ans := readNum(reader)
	if res != ans {
		t.Fatalf("Sample expect %d, but got %d", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `9 4
aaaaaaaaa
aaaa
5`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `7 3
abacaba
aaa
6`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `5 4
aabcd
abcd
3`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `4 2
abba
bb
4`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `6 4
baraka
baka
4`
	runSample(t, s)
}

func TestSample6(t *testing.T) {
	s := `8 7
question
problem
-1`
	runSample(t, s)
}

func TestSample7(t *testing.T) {
	s := `6 4
bcccbc
cccc
5`
	runSample(t, s)
}

func TestSample8(t *testing.T) {
	s := `6 4
ccbaca
cbaa
5`
	runSample(t, s)
}

func TestSample9(t *testing.T) {
	s := `5 4
aacde
acde
3`
	runSample(t, s)
}

func TestSample10(t *testing.T) {
	s := `7 5
caacbaa
cacba
5`
	runSample(t, s)
}

func TestSample11(t *testing.T) {
	s := `6 5
ababab
aabab
4`
	runSample(t, s)
}
