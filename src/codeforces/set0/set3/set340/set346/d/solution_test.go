package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, s1, s2, virus := process(reader)
	expect := readString(reader)
	if res == expect {
		return
	}
	if len(res) != len(expect) || res == "0" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	// res is sub-sequence of s1 and s2
	checkSubSequence := func(x string, y string) bool {
		for i, j := 0, 0; i < len(y); i++ {
			for j < len(x) && x[j] != y[i] {
				j++
			}
			if j == len(x) {
				return false
			}
			j++
		}
		return true
	}

	if !checkSubSequence(s1, res) {
		t.Fatalf("Sample result %s, is not a sub-sequence of %s", res, s1)
	}

	if !checkSubSequence(s2, res) {
		t.Fatalf("Sample result %s, is not a sub-sequence of %s", res, s2)
	}

	// virus not a substring of res
	if strings.Contains(res, virus) {
		t.Fatalf("Sample result %s, is a substring of %s", virus, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `AJKEQSLOBSROFGZ
OVGURWZLWVLUXTH
OZ
ORZ`)
}

func TestSample2(t *testing.T) {
	runSample(t, `ABC
BC
C
B`)
}

func TestSample3(t *testing.T) {
	runSample(t, `AA
A
A
0`)
}

func TestSample4(t *testing.T) {
	runSample(t, `ABBB
ABBB
ABB
BBB`)
}

func TestSample5(t *testing.T) {
	runSample(t, `DASSDASDASDDAASDASDADASDASASDAS
SDADASDASSDAASDASDASDADASSDDA
SD
DADADADAADADADADASSA`)
}
