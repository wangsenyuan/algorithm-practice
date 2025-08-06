package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	cnt, best := process(reader)
	var expect_cnt int
	fmt.Fscanf(reader, "%d\n", &expect_cnt)
	expect_best := readString(reader)

	if cnt != expect_cnt || best != expect_best {
		t.Fatalf("Sample expect %d, %s, but got %d, %s", expect_cnt, expect_best, cnt, best)
	}
}

func TestSample1(t *testing.T) {
	s := `ABA
CBA
1
ABC`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `CDBABC
ADCABD
2
ADBADC`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `OVGHK
RPGUC
4
CPGRU`
	runSample(t, s)
}
