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
	ss := strings.Split(s, " ")
	k, _ := strconv.Atoi(ss[1])
	s = readString(reader)
	return solve(k, s)
}

const mod = 1e9 + 7

func add(nums ...int) int {
	var res int
	for _, num := range nums {
		res += num
		if res >= mod {
			res -= mod
		}
	}
	return res
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res = res * num % mod
	}
	return res
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = r * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, mod-2)
}

func solve(k int, s string) int {
	n := len(s)
	pw := make([]int, n+1)
	pw[0] = 1

	F := make([]int, n+1)
	F[0] = 1

	for i := range n {
		pw[i+1] = mul(pw[i], 10)
		F[i+1] = mul(F[i], i+1)
	}

	I := make([]int, n+1)
	I[n] = inverse(F[n])
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], I[r], I[n-r])
	}

	// when l = 0
	var sum int
	for i := range n {
		sum = add(sum, int(s[i]-'0'))
	}

	var res int

	for l := 0; l <= n-2; l++ {
		sum = sub(sum, int(s[n-1-l]-'0'))
		res = add(res, mul(pw[l], nCr(n-l-2, k-1), sum))
	}

	for i := range n {
		x := int(s[i] - '0')
		tmp := mul(x, pw[n-1-i], nCr(i, k))
		res = add(res, tmp)
	}

	return res
}
