package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

const mod = 1e9 + 7

func solve(x string) int {
	n := len(x)
	pw := make([]int, n+1)
	pw[0] = 1
	for i := 1; i <= n; i++ {
		pw[i] = pw[i-1] * 2 % mod
	}

	var res int
	for i := range n {
		if x[i] == '0' {
			continue
		}
		// 每一对 0??? 1??? 都会贡献一次
		res += pw[n-i-1] * pw[n-1] % mod
		res %= mod
	}
	return res
}
