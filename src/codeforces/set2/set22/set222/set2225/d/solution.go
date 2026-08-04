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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n, x int64
	fmt.Fscan(reader, &n, &x)
	return solve(n, x)
}

func solve(n, x int64) int64 {
	leftZero := countPrefixValueZero(0, x-1)
	rightZero := countPrefixValueZero(x, n)
	leftOne := countMod(0, x-1, 1)
	rightOne := countMod(x, n, 1)

	return (mulMod(leftZero, rightZero) + mulMod(leftOne, rightOne)) % mod
}

const mod int64 = 998244353

func mulMod(a, b int64) int64 {
	return (a % mod) * (b % mod) % mod
}

func countPrefixValueZero(l, r int64) int64 {
	if l > r {
		return 0
	}
	res := countMod(l, r, 3)
	if l == 0 {
		res++
	}
	return res
}

func countMod(l, r int64, rem int64) int64 {
	if l > r {
		return 0
	}
	return countModPrefix(r, rem) - countModPrefix(l-1, rem)
}

func countModPrefix(n int64, rem int64) int64 {
	if n < 0 {
		return 0
	}
	if n < rem {
		return 0
	}
	return (n-rem)/4 + 1
}
