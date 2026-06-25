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
	a, res := drive(reader)

	calc := func(arr []int) int {
		var res int
		for i := range len(arr) {
			sum := a[i]
			for j := i + 1; j < len(arr); j++ {
				if arr[i] > arr[j] {
					res += sum
				}
				sum += a[j]
			}
		}
		return res
	}

	x := calc(expect)
	y := calc(res)

	if x != y {
		t.Fatalf("Sample expect %v(%d), but got %v(%d)", expect, x, res, y)
	}

}

func TestSample1(t *testing.T) {
	runSample(t, `1
1
0
`, []int{1})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2
1000000000 -1000000000
`, []int{2, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
3
1 2 3
`, []int{3, 2, 1})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
4
-1 -2 -3 -4
`, []int{1, 2, 3, 4})
}

func TestSample5(t *testing.T) {
	runSample(t, `1
5
-1 2 -3 2 -1
`, []int{3, 4, 1, 5, 2})
}

func TestSample6(t *testing.T) {
	runSample(t, `1
6
1 -1 3 -4 1 -3
`, []int{5, 2, 4, 1, 6, 3})
}

func TestSample7(t *testing.T) {
	runSample(t, `1
7
-3 -2 -1 4 -1 -2 -3
`, []int{1, 4, 6, 7, 2, 3, 5})
}
