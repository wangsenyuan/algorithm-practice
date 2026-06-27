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
3 2
010
110
`, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
1 1
0
0
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
2 2
01
00
`, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
7 30
1010111
0011010
`, 12169074016)
}

func bruteSolve(n, k int, s, z string) int64 {
	a := []string{s, z}
	for step := 0; step < k; step++ {
		b := make([]string, 0, len(a)*2-1)
		for i := 0; i+1 < len(a); i++ {
			b = append(b, a[i])
			buf := make([]byte, n)
			for j := 0; j < n; j++ {
				if a[i][j] == a[i+1][j] {
					buf[j] = '0'
				} else {
					buf[j] = '1'
				}
			}
			b = append(b, string(buf))
		}
		b = append(b, a[len(a)-1])
		a = b
	}
	var res int64
	for _, cur := range a {
		var cnt int64
		for i := 0; i < n; i++ {
			cnt += int64(cur[i] - '0')
		}
		res += cnt * int64(n-int(cnt))
	}
	return res
}

func TestSmallAgainstBruteForce(t *testing.T) {
	for n := 1; n <= 5; n++ {
		limit := 1 << n
		for k := 1; k <= 4; k++ {
			for x := 0; x < limit; x++ {
				for y := 0; y < limit; y++ {
					s := fmt.Sprintf("%0*b", n, x)
					z := fmt.Sprintf("%0*b", n, y)
					expect := bruteSolve(n, k, s, z)
					res := solve(n, k, s, z)
					if res != expect {
						t.Fatalf("n=%d k=%d s=%s z=%s expect %d, but got %d", n, k, s, z, expect, res)
					}
				}
			}
		}
	}
}
