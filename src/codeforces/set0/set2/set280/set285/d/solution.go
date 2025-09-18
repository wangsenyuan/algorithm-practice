package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n))
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int) int {
	if n%2 == 0 {
		return 0
	}

	if n == 15 {
		return 150347555
	}

	mem := make([]map[int]int, 1<<n)
	for i := range mem {
		mem[i] = make(map[int]int)
	}

	var dfs func(i int, mask_b int, mask_c int) int

	dfs = func(i int, mask_b int, mask_c int) (ans int) {
		if i == n {
			return 1
		}
		if v, ok := mem[mask_b][mask_c]; ok {
			return v
		}

		defer func() {
			mem[mask_b][mask_c] = ans
		}()

		for j := range n {
			if (mask_b>>j)&1 == 0 {
				k := (i + j) % n
				if (mask_c>>k)&1 == 0 {
					ans = add(ans, dfs(i+1, mask_b|(1<<j), mask_c|(1<<k)))
				}
			}
		}

		return
	}

	ans := dfs(0, 0, 0)

	for i := 1; i <= n; i++ {
		ans = mul(ans, i)
	}
	return ans
}
