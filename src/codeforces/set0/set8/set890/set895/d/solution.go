package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	a := readString(reader)
	b := readString(reader)
	return solve(a, b)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, mod-2)
}

func div(a, b int) int {
	return mul(a, inverse(b))
}

func solve(a string, b string) int {
	n := len(a)
	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}
	I := make([]int, n+1)
	I[n] = pow(F[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	freq := make([]int, 26)
	for i := range n {
		freq[int(a[i]-'a')]++
	}

	suf := make([]int, 27)

	// 找到a的一个排列 < s
	play := func(s string) int {
		f := slices.Clone(freq)

		var res int

		for i := range n {
			x := int(s[i] - 'a')
			// 先计算进来
			suf[26] = 1
			// (m - i - 1)! / (f[1]! * f[2]!... * f[25]!)
			for c := 25; c >= 0; c-- {
				suf[c] = mul(suf[c+1], I[f[c]])
			}

			w := F[n-i-1]
			pref := 1

			for c := range x {
				if f[c] > 0 {
					tmp := mul(pref, suf[c+1])
					tmp = mul(tmp, I[f[c]-1])
					res = add(res, mul(w, tmp))
				}
				pref = mul(pref, I[f[c]])
			}

			if f[x] == 0 {
				break
			}
			f[x]--
		}
		return res
	}

	res := play(b)
	res = sub(res, play(a))
	// 去掉a
	res = sub(res, 1)
	return res
}
