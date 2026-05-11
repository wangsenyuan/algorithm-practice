package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
5 5 2 3 0
3 0 4 1
`, []int{3, 3, 4, 2, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `6
4 3 0 3 0 2
1 0 0 1 3
`, []int{1, 0, 3, 3, 2, 3})
}

func TestSmallAgainstBruteForce(t *testing.T) {
	for n := 2; n <= 4; n++ {
		totalA := 1
		for i := 0; i < n; i++ {
			totalA *= n + 1
		}
		totalH := 1
		for i := 0; i < n-1; i++ {
			totalH *= n
		}
		for maskA := 0; maskA < totalA; maskA++ {
			a := decode(maskA, n, n+1)
			for maskH := 0; maskH < totalH; maskH++ {
				h := decode(maskH, n-1, n)
				expect := bruteForce(a, h)
				res := solve(a, h)
				if !reflect.DeepEqual(res, expect) {
					t.Fatalf("n=%d a=%v h=%v expect %v, but got %v", n, a, h, expect, res)
				}
			}
		}
	}
}

func decode(mask int, n int, base int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = mask % base
		mask /= base
	}
	return res
}

func bruteForce(a []int, h []int) []int {
	n := len(a)
	res := make([]int, n)
	for y := 1; y <= n; y++ {
		row := make([]bool, n)
		for i := 0; i < n; i++ {
			row[i] = a[i] >= y
		}
		for l := 0; l < n; {
			r := l
			for r < n-1 && h[r] < y {
				r++
			}
			var cnt int
			for i := l; i <= r; i++ {
				if row[i] {
					cnt++
				}
			}
			for i := r - cnt + 1; i <= r; i++ {
				res[i]++
			}
			l = r + 1
		}
	}
	return res
}
