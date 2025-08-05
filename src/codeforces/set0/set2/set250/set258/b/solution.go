package main

import "fmt"

func main() {
	var m int
	fmt.Scanf("%d", &m)
	fmt.Println(solve(m))
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

func solve(m int) int {
	digits := fmt.Sprintf("%d", m)

	n := len(digits)

	var f [15][15][2]int

	f[0][0][1] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= 9; k++ {
				nj := j
				if k == 4 || k == 7 {
					nj++
				}
				f[i][nj][0] = add(f[i][nj][0], f[i-1][j][0])
			}
			for k := 0; k <= int(digits[i-1]-'0'); k++ {
				nj := j
				if k == 4 || k == 7 {
					nj++
				}
				nw := 0
				if k == int(digits[i-1]-'0') {
					nw = 1
				}
				f[i][nj][nw] = add(f[i][nj][nw], f[i-1][j][1])
			}
		}
	}

	f[n][0][0] = sub(f[n][0][0], 1)

	var ans int

	var dfs func(x int, c int, s int)
	dfs = func(x int, c int, s int) {
		if x > 6 {
			for i := c + 1; i <= n; i++ {
				ans = add(ans, mul(s, add(f[n][i][0], f[n][i][1])))
			}
			return
		}
		if c >= n {
			return
		}
		for i := 0; i <= n; i++ {
			if f[n][i][0]+f[n][i][1] > 0 {
				f[n][i][0] = sub(f[n][i][0], 1)
				dfs(x+1, c+i, s*(f[n][i][0]+f[n][i][1]+1)%mod)
				f[n][i][1] = add(f[n][i][1], 1)
			}
		}
	}

	dfs(1, 0, 1)
	return ans
}
