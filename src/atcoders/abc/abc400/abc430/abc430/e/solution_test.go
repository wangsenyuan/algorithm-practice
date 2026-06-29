package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	var tc int
	fmt.Fscan(reader, &tc)
	for i := range tc {
		res := drive(reader)
		if res != expect[i] {
			t.Fatalf("Sample case %d expect %d, but got %d", i+1, expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1010001
1000110
000
111
01010
01010
0101
0011
100001101110000001010110110001
101100011000011011100000010101
`, []int{2, -1, 0, -1, 22})
}
