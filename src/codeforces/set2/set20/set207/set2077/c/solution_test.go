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
	s := `3
3 2
010
1
3
10 3
0101000110
3
5
10
24 1
011001100110000101111000
24
`
	expect := []int{1, 5, 512, 768, 1536, 23068672}
	runSample(t, s, expect)
}

func TestSingleCharacter(t *testing.T) {
	s := `2
1 2
0
1
1
1 1
1
1
`
	expect := []int{0, 0, 0}
	runSample(t, s, expect)
}

func TestFormulaAgainstBruteForce(t *testing.T) {
	for n := 1; n <= 8; n++ {
		base := pow(2, n) * pow(16, mod-2) % mod
		for mask := 0; mask < 1<<n; mask++ {
			cnt0 := 0
			buf := make([]byte, n)
			for i := 0; i < n; i++ {
				if mask>>i&1 == 1 {
					buf[i] = '1'
				} else {
					buf[i] = '0'
					cnt0++
				}
			}
			expect := bruteScore(string(buf))
			res := calc(n, cnt0, base)
			if res != expect {
				t.Fatalf("String %s expect %d, but got %d", string(buf), expect, res)
			}
		}
	}
}

func bruteScore(s string) int {
	n := len(s)
	var res int
	for mask := 1; mask < 1<<n; mask++ {
		var vals []int
		for i := 0; i < n; i++ {
			if mask>>i&1 == 1 {
				if s[i] == '1' {
					vals = append(vals, 1)
				} else {
					vals = append(vals, -1)
				}
			}
		}
		best := 0
		for i := 0; i <= len(vals); i++ {
			left, right := 0, 0
			for j := 0; j < i; j++ {
				left += vals[j]
			}
			for j := i; j < len(vals); j++ {
				right += vals[j]
			}
			best = max(best, left*right)
		}
		res += best
	}
	return res % mod
}
