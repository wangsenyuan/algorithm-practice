package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `01234567
34576210`
	expect := []int{2, 6}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `Die Polizei untersucht eine Straftat im IT-Bereich.
untersucht eine Straftat.hciereB-TI mi  ieziloP eiD`
	expect := []int{11, 36}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `cbaaaa
aaaabc`
	expect := []int{4, 5}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `123342
3324212
`
	expect := []int{-1, -1}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `q
q
`
	expect := []int{-1, -1}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `:(((((((((((
(((((((((((%
`
	expect := []int{-1, -1}
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := ` 987654321  -  0001111102 
54321  -  0001111 2016789 
`
	expect := []int{4, 22}
	runSample(t, s, expect)
}
