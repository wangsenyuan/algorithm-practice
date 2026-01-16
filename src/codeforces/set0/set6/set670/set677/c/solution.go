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

const mod = 1_000_000_007

func mul(a, b int) int {
	return a * b % mod
}

func id(x byte) int {
	if x >= '0' && x <= '9' {
		return int(x - '0')
	}
	if x >= 'A' && x <= 'Z' {
		return 10 + int(x-'A')
	}
	if x >= 'a' && x <= 'z' {
		return 36 + int(x-'a')
	}
	if x == '-' {
		return 62
	}
	return 63
}

func solve(s string) int {
	var dp [64]int

	for i := range 64 {
		for j := range 64 {
			dp[i&j]++
		}
	}

	res := 1
	for i := range len(s) {
		x := id(s[i])
		res = mul(res, dp[x])
	}

	return res
}
