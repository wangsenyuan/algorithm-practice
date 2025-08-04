package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	var expect int
	fmt.Fscan(reader, &expect)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2
2 1
1 2.0
1 3.100
1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
1 5.0
2 5.5
3 6.0
0`)
}

func TestSample3(t *testing.T) {
	runSample(t, `6 3
1 14.284235
2 17.921382
1 20.328172
3 20.842331
1 25.790145
1 27.204125
2`)
}

func TestSample4(t *testing.T) {
	runSample(t, `15 5
4 6.039627
2 7.255149
2 14.469785
2 15.108572
4 22.570081
5 26.642253
5 32.129202
5 44.288220
5 53.231909
5 60.548042
4 62.386581
2 77.828816
1 87.998512
3 96.163559
2 99.412872
6`)
}
