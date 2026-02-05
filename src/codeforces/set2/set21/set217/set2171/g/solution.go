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
		fmt.Fprintln(writer, res[0], res[1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

const mod = int(1e6 + 3)

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

var F [mod]int
var I [mod]int

func init() {
	F[0] = 1
	for i := 1; i < mod; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[mod-1] = pow(F[mod-1], mod-2)
	for i := mod - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func solve(a []int, b []int) []int {
	// 需要多少次double
	// 如果全部是偶数，且能够double的时候，就double
	ans := 1
	var x int

	n := len(a)
	// 对于单个数来说，它如果要double，那么这个double最好是放在越后面进行
	d := 30

	for i, v := range a {
		w := b[i]
		var tmp int
		for w/2 >= v {
			tmp++
			w /= 2
		}
		d = min(d, tmp)
	}

	for range d {
		// 这个阶段，先把w-1的处理掉
		var cnt int
		for i := range n {
			w := b[i]
			if w&1 == 1 {
				cnt++
			}
			b[i] /= 2
		}
		x += cnt
		ans = mul(ans, F[cnt])
		x++
	}

	diff := make([]int, n)
	var sum int
	for i := range n {
		diff[i] = b[i] - a[i]
		sum += diff[i]
	}
	x += sum
	if sum >= mod {
		ans = 0
	} else {
		ans = mul(ans, F[sum])
		for i := range n {
			ans = mul(ans, I[diff[i]])
		}
	}

	return []int{x, ans}
}
