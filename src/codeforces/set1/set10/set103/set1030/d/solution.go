package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if !res.ok {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	for _, p := range res.pts {
		fmt.Println(p[0], p[1])
	}
}

type result struct {
	ok  bool
	pts [][]int
}

func drive(reader *bufio.Reader) result {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	return solve(n, m, k)
}

func solve(n, m, k int) result {
	g := gcd(n*m, k)
	k1 := k / g
	if 2%k1 != 0 {
		// k1 = 2 or k1 = 1
		return result{ok: false}
	}
	// w := n * m / g * 2 / k1
	// a * b = w

	play := func(a int, b int) [][]int {
		return [][]int{{0, 0}, {a, 0}, {0, b}}
	}

	for i := 1; i <= g/i; i++ {
		if g%i == 0 {
			j := g / i
			if 2/k1*n%i == 0 && m%j == 0 {
				a := 2 / k1 * n / i
				b := m / j
				if a <= n && b <= m {
					return result{ok: true, pts: play(a, b)}
				}
				if a <= m && b <= n {
					return result{ok: true, pts: play(b, a)}
				}
			}
			if n%i == 0 && 2/k1*m%j == 0 {
				a := n / i
				b := 2 / k1 * m / j
				if a <= n && b <= m {
					return result{ok: true, pts: play(a, b)}
				}
				if a <= m && b <= n {
					return result{ok: true, pts: play(b, a)}
				}
			}
		}
	}
	return result{ok: false}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
