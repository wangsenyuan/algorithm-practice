package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `50 50 7
10 F 1 80000
20 T 1 80000
30 T 81000 82000
40 T 83000 84000
50 T 85000 86000
60 T 87000 88000
70 F 81000 89000
`
	runSample(t, s, 100)
}

func TestSample2(t *testing.T) {
	s := `80 72 9
15 T 8210 15679
10 F 11940 22399
50 T 30600 44789
50 F 32090 36579
5 F 45520 48519
120 F 49250 55229
8 F 59700 80609
35 T 61940 64939
2 T 92540 97769
`
	runSample(t, s, 120)
}