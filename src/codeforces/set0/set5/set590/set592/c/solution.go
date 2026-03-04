package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%d/%d\n", res[0], res[1])
}

func drive(reader *bufio.Reader) []int {
	var t, w, b int
	fmt.Fscan(reader, &t, &w, &b)
	return solve(t, w, b)
}

func solve(t int, w int, b int) []int {
	if w == b {
		return []int{1, 1}
	}

	tmp := lcm(w, b)
	var cnt int
	if tmp.Cmp(big.NewInt(int64(t))) > 0 {
		cnt = 1
	} else {
		cnt = t/int(tmp.Int64()) + 1
	}
	dist := min(w, b)
	// cnt := t/x + 1
	// 0 不算
	tot := dist*(cnt-1) - 1
	last := int(tmp.Int64()) * (cnt - 1)
	tot += min(t+1, last+dist) - last
	// 计算最后
	c := gcd(tot, t)

	return []int{tot / c, t / c}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) *big.Int {
	c := gcd(a, b)

	res := big.NewInt(int64(a / c))

	res = res.Mul(res, big.NewInt(int64(b)))

	return res
}
