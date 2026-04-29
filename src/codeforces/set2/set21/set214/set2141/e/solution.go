package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var s string
		fmt.Fscan(reader, &s)
		res := solve(s)
		fmt.Fprintln(writer, res)
	}
}

const mod = 998244353

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

func solve(s string) int {
	// n := len(s)
	res := 1
	for _, v := range s {
		if v == '?' {
			res = mul(res, 2)
		}
	}
	// 只有0111不可以
	if check(s) {
		res = sub(res, 1)
	}
	return res
}

func check(s string) bool {
	if s[0] == '1' {
		return false
	}
	// s[0] == '0' or '?'
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			return false
		}
	}
	return true
}
