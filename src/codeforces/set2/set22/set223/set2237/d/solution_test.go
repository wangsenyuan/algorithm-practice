package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
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
1
0
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2
01
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
5
01001
`, 10)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
3
001
`, 5)
}

func TestSample5(t *testing.T) {
	runSample(t, `1
6
011110
`, 15)
}

func TestSample6(t *testing.T) {
	runSample(t, `1
9
010110110
`, 30)
}

func TestSample7(t *testing.T) {
	runSample(t, `1
12
010000101001
`, 47)
}

func TestSample8(t *testing.T) {
	runSample(t, `1
16
1010011010010110
`, 81)
}

func TestSample9(t *testing.T) {
	runSample(t, `1
20
11110101101101001110
`, 139)
}

func TestSample10(t *testing.T) {
	runSample(t, `1
30
000101100011111001111100000010
`, 316)
}
