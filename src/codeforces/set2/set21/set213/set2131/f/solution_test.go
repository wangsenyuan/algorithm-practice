package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
11
00
`
	runSample(t, s, 5)
}

func TestSample2(t *testing.T) {
	s := `2
01
01
`
	runSample(t, s, 4)
}

func TestSample3(t *testing.T) {
	s := `4
1010
1101
`
	runSample(t, s, 24)
}

func TestSmallAgainstBruteForce(t *testing.T) {
	for n := 1; n <= 5; n++ {
		for maskA := 0; maskA < 1<<n; maskA++ {
			for maskB := 0; maskB < 1<<n; maskB++ {
				a := makeBinaryString(maskA, n)
				b := makeBinaryString(maskB, n)
				res := solve(a, b)
				expect := bruteForce(a, b)
				if res != expect {
					t.Fatalf("n=%d a=%s b=%s expect %d, but got %d", n, a, b, expect, res)
				}
			}
		}
	}
}

func makeBinaryString(mask int, n int) string {
	buf := make([]byte, n)
	for i := range n {
		if mask>>i&1 == 1 {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}
	return string(buf)
}

func bruteForce(a string, b string) int64 {
	n := len(a)
	var ans int64
	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			best := 2 * n
			for maskA := 0; maskA < 1<<n; maskA++ {
				for maskB := 0; maskB < 1<<n; maskB++ {
					cost := 0
					aa := []byte(a)
					bb := []byte(b)
					for i := range n {
						if maskA>>i&1 == 1 {
							aa[i] ^= 1
							cost++
						}
						if maskB>>i&1 == 1 {
							bb[i] ^= 1
							cost++
						}
					}
					if cost < best && reachable(aa, bb, x, y) {
						best = cost
					}
				}
			}
			ans += int64(best)
		}
	}
	return ans
}

func reachable(a []byte, b []byte, x int, y int) bool {
	dp := make([][]bool, x)
	for i := range x {
		dp[i] = make([]bool, y)
	}
	for i := range x {
		for j := range y {
			if a[i]^b[j] != 0 {
				continue
			}
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i > 0 && dp[i-1][j] {
				dp[i][j] = true
			} else if j > 0 && dp[i][j-1] {
				dp[i][j] = true
			}
		}
	}
	return dp[x-1][y-1]
}
