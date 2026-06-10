package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var tc int
	fmt.Fscan(reader, &tc)
	var res []int
	for ; tc > 0; tc-- {
		var n, q int
		var s string
		fmt.Fscan(reader, &n, &q, &s)
		queries := make([]int, q)
		for i := range queries {
			fmt.Fscan(reader, &queries[i])
		}
		res = append(res, solve(s, queries)...)
	}
	return res
}

func solve(s string, queries []int) []int {
	buf := []byte(s)
	n := len(buf)
	var cnt0 int
	for _, x := range buf {
		if x == '0' {
			cnt0++
		}
	}

	base := pow(2, n) * pow(16, mod-2) % mod
	ans := make([]int, len(queries))

	for i, pos := range queries {
		pos--
		if buf[pos] == '0' {
			buf[pos] = '1'
			cnt0--
		} else {
			buf[pos] = '0'
			cnt0++
		}
		ans[i] = calc(n, cnt0, base)
	}

	return ans
}

func calc(n int, cnt0 int, base int) int {
	x := (int64(n)*(int64(n)+1) - 4*int64(cnt0)*int64(n) + 4*int64(cnt0)*int64(cnt0) - 2) % mod
	if x < 0 {
		x += mod
	}
	return int(int64(base) * x % mod)
}

func pow(a int, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = int(int64(res) * int64(a) % mod)
		}
		a = int(int64(a) * int64(a) % mod)
		n >>= 1
	}
	return res
}
