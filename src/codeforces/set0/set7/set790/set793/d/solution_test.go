package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `7 4
4
1 6 2
6 2 2
2 4 2
2 7 1
	`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 3
4
2 1 2
1 3 2
3 4 2
4 1 1
	`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 2
10
2 3 290
3 1 859
3 1 852
1 2 232
1 2 358
2 1 123
1 3 909
2 1 296
1 3 119
1 2 584
	`, 119)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 5
20
2 5 174
4 3 496
5 2 103
2 1 345
2 4 942
3 5 131
3 2 451
5 2 299
2 4 285
4 5 241
4 5 706
2 1 639
1 5 94
1 2 844
3 4 194
2 4 812
2 5 566
3 5 293
3 4 356
2 5 717
	`, 978)
}
