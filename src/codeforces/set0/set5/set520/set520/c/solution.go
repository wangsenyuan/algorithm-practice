package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

const mod = 1000000007

func mul(a, b int) int {
	return (a * b) % mod
}
func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func solve(s string) int {
	freq := make([]int, 4)
	n := len(s)
	for i := range n {
		switch s[i] {
		case 'A':
			freq[0]++
		case 'C':
			freq[1]++
		case 'G':
			freq[2]++
		case 'T':
			freq[3]++
		}
	}

	x := slices.Max(freq)
	var res int
	for _, v := range freq {
		if v == x {
			res++
		}
	}

	return pow(res, n)
}
