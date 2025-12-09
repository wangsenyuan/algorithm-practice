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

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const mod = 998244353

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

func solve(a []int) int {
	pos := make(map[int]int)
	for i, v := range a {
		pos[v] = i
	}

	var cnt int

	n := len(a)

	for i := 0; i < n; {
		cnt++
		if pos[a[i]] == i {
			i++
			continue
		}
		r := pos[a[i]]
		for j := i; j <= r; j++ {
			r = max(r, pos[a[j]])
		}
		// 这一段的数字必须相同
		i = r + 1
	}
	return pow(2, cnt-1)
}
