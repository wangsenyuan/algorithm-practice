package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	var tc int
	fmt.Fscan(reader, &tc)
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
3
400
500
200
100
300
600
`, 1200)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
6
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
`, 6000000000)
}
