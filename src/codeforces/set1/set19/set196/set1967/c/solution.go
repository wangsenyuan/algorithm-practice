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
		res := drive(reader)
		for i, v := range res {
			if i+1 < len(res) {
				fmt.Fprintf(writer, "%d ", v)
			} else {
				fmt.Fprintln(writer, v)
			}
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

const mod = 998244353

func mul(a, b int) int {
	return a * b % mod
}

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

const N = 1e6 + 100

var inv [N]int

func init() {

	inv[0] = 1
	inv[1] = 1
	for i := 2; i < N; i++ {
		inv[i] = mul(mod-mod/i, inv[mod%i])
	}
}

func lowbit(x int) int {
	return x & -x
}

func solve(k int, a []int) []int {
	n := len(a)

	for i := range n {
		tmp := 1
		d := 1
		for u := i + 1 + lowbit(i+1); u <= n; u += lowbit(u) {
			tmp = mul(tmp, mul(d+k-1, inv[d]))
			a[u-1] = sub(a[u-1], mul(a[i], tmp))
			d++
		}
	}

	return a
}
