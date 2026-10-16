package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 5 5 6
4 3
1 3
3 3
2 5
3 2
1 3
1 4
2 3
2 4
1 1
1 1
`, []string{
		"Waste",
		"Grapes",
		"Carrots",
		"Kiwis",
		"Carrots",
		"Carrots",
	})
}

func TestSample2(t *testing.T) {
	runSample(t, `2 3 2 2
1 1
2 2
2 1
2 2
`, []string{
		"Grapes",
		"Waste",
	})
}

func TestSample3(t *testing.T) {
	runSample(t, `31 31 31 4
4 9
16 27
11 29
8 28
11 2
10 7
22 6
1 25
14 8
9 7
9 1
2 3
5 2
21 16
20 19
23 14
27 6
25 21
14 1
18 14
7 2
19 12
30 27
4 27
24 12
25 20
26 22
21 17
11 6
5 28
28 24
17 30
2 5
30 10
4 21
`, []string{
		"Kiwis",
		"Carrots",
		"Grapes",
		"Grapes",
	})
}
