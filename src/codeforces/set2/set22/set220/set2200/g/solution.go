package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	ops := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &ops[i])
	}
	return solve(x, ops)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

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

func inverse(num int) int {
	return pow(num, mod-2)
}

func solve(x int, ops []string) int {
	var num int
	var muls []int
	prod := 1
	for _, op := range ops {
		y, _ := strconv.Atoi(op[1:])

		switch op[0] {
		case '+':
			num = add(num, y)
		case '-':
			num = add(num, mod-y)
		case 'x':
			muls = append(muls, y)
			prod = mul(prod, y)
		default:
			// div
			muls = append(muls, inverse(y))
			prod = mul(prod, inverse(y))
		}
	}

	// dp[k] is the sum of products over all size-k subsets of muls.
	dp := make([]int, len(muls)+1)
	dp[0] = 1
	for i, y := range muls {
		for k := i + 1; k > 0; k-- {
			dp[k] = add(dp[k], mul(dp[k-1], y))
		}
	}

	// For a fixed addition, a particular size-k subset is applied after it
	// with probability k! * (m-k)! / (m+1)!.
	m := len(muls)
	fac := make([]int, m+1)
	fac[0] = 1
	for i := 1; i <= m; i++ {
		fac[i] = mul(fac[i-1], i)
	}
	invFacM := inverse(fac[m])
	expectMultiplier := 0
	for k := 0; k <= m; k++ {
		weight := mul(mul(fac[k], fac[m-k]), invFacM)
		expectMultiplier = add(expectMultiplier, mul(dp[k], weight))
	}
	expectMultiplier = mul(expectMultiplier, inverse(m+1))

	return add(mul(x, prod), mul(num, expectMultiplier))
}
