package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	s := readString(reader)
	s1 := readString(reader)
	k, _ := strconv.Atoi(s1)
	return solve(s, k)
}

const mod = 1_000_000_007

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

func inverse(a int) int {
	return pow(a, mod-2)
}

func solve(s string, k int) int {
	m := len(s)
	// sum = (q^^n - 1) / (q - 1)
	// q = 2 ^^ m
	q := pow(2, m)
	q2 := sub(pow(q, k), 1)
	q1 := inverse(sub(q, 1))
	var sum int
	for i := range m {
		if s[i] == '0' || s[i] == '5' {
			sum = add(sum, pow(2, i))
		}
	}
	return mul(sum, mul(q2, q1))
}
