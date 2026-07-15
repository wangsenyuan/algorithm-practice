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
	a, b := n, m

	g := gcd(a, k)
	a /= g
	k /= g

	g = gcd(b, k)
	b /= g
	k /= g

	if k > 2 {
		return result{ok: false}
	}

	if k == 1 {
		if 2*a <= n {
			a *= 2
		} else {
			b *= 2
		}
	}

	pts := [][]int{{0, 0}, {a, 0}, {0, b}}
	return result{ok: true, pts: pts}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
