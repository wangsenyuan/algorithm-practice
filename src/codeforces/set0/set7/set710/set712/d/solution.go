package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

const mod = 1000000007
const max = 1000005

var f [max]int
var fi [max]int

func pos(a int) int {
	return ((a%mod + mod) % mod)
}

func add(a, b int) int {
	return ((a + b) % mod)
}

func sub(a, b int) int {
	return pos(a - b)
}

func mult(a, b int) int {
	return pos(a) * b % mod
}

func sq(a int) int {
	return int(int64(a) * int64(a) % mod)
}

func expo(a, b int) int {
	if b == 0 {
		return 1
	}
	if b%2 == 1 {
		return mult(a, sq(expo(a, b/2)))
	}
	return sq(expo(a, b/2))
}

func inv(a int) int {
	return expo(a, mod-2)
}

func c(n, k int) int {
	if n < 0 || k < 0 || k > n || n >= max {
		return 0
	}
	return mult(f[n], mult(fi[k], fi[n-k]))
}

func drive(reader *bufio.Reader) int {
	var a, b, k, t int
	fmt.Fscan(reader, &a, &b, &k, &t)
	return solve(a, b, k, t)
}

func solve(a int, b int, k int, t int) int {
	// Initialize factorials and inverse factorials
	f[0] = 1
	fi[0] = 1
	for i := 1; i < max; i++ {
		f[i] = mult(f[i-1], i)
		fi[i] = inv(f[i])
	}

	// Early exit condition
	if b-a > 2*k*t {
		return 0
	}

	// Initialize poly1 array
	poly1 := make([]int, max)
	for i := 0; i <= 2*t; i++ {
		sign := 1
		if i%2 != 0 {
			sign = -1
		}
		poly1[(2*k+1)*i] = sign * c(2*t, i)
		// Normalize to positive modulo
		if poly1[(2*k+1)*i] < 0 {
			poly1[(2*k+1)*i] = pos(poly1[(2*k+1)*i])
		}
	}

	// Initialize pref2 array
	// We only need up to ub = 4*k*t, but compute a bit more for safety
	maxIdx := 4*k*t + 1
	if maxIdx >= max {
		maxIdx = max - 1
	}
	pref2 := make([]int, maxIdx+1)
	for i := 0; i <= maxIdx; i++ {
		n := 2*t - 1 + i + 1
		if n >= max {
			break
		}
		pref2[i] = c(n, 2*t)
	}

	// Calculate bounds
	lb := 2*k*t + b - a + 1
	ub := 4 * k * t

	var ans int
	for i := 0; i <= 2*t; i++ {
		l := lb - (2*k+1)*i
		u := ub - (2*k+1)*i
		if u < 0 {
			break
		}
		if u >= len(pref2) {
			continue
		}
		if l < 0 {
			ans = add(ans, mult(poly1[(2*k+1)*i], pref2[u]))
		} else {
			if l-1 >= 0 && l-1 < len(pref2) {
				ans = add(ans, mult(poly1[(2*k+1)*i], sub(pref2[u], pref2[l-1])))
			} else {
				// l-1 < 0, so pref2[l-1] = 0 (prefix sum before index 0)
				ans = add(ans, mult(poly1[(2*k+1)*i], pref2[u]))
			}
		}
	}

	return ans
}
