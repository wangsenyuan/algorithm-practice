package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 6
1 3 5 4
3 2 3
1 1 6 2
2 3 4
3 1 6
3 2 3
`, []int64{4, 6, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, `2 8
1 1 2 1000000000
1 1 2 1000000000
2 2 2
1 1 2 1000000000
1 1 2 1000000000
1 1 2 1000000000
3 2 2
3 1 2
`, []int64{0, 5000000000})
}

func TestSample3(t *testing.T) {
	runSample(t, `24 30
1 11 24 4326
1 4 16 1149
1 14 20 2331
1 12 14 8930
1 22 23 6989
3 15 20
3 10 19
1 3 12 7988
1 18 23 8450
3 9 19
3 13 15
2 8 15
2 9 14
1 11 17 4062
1 6 15 1721
3 7 13
1 11 20 8541
1 8 10 3748
1 1 17 3252
2 9 23
2 1 23
3 2 22
1 5 23 7468
3 1 12
3 12 19
2 6 24
3 2 14
3 1 15
2 15 19
3 2 14
`, []int64{7806, 16736, 22393, 16736, 10858, 0, 7468, 7468, 0, 0, 0})
}
