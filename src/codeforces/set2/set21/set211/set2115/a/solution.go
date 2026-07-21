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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
	for i := range t {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[j])
		}
		res[i] = solve(a)
	}
	return res
}

const X = 5010

var g [X][X]int

func init() {
	for i := range X {
		g[i][0] = i
		g[0][i] = i
		g[i][i] = i
	}

	for i := 1; i < X; i++ {
		for j := 1; j < i; j++ {
			g[i][j] = g[j][i%j]
			g[j][i] = g[i][j]
		}
	}
}

func solve(a []int) int {
	var k int
	for _, v := range a {
		k = g[k][v]
	}

	var mx int
	for _, v := range a {
		v /= k
		mx = max(mx, v)
	}

	f := make([]int, mx+1)
	for i := range f {
		f[i] = inf
	}

	for _, v := range a {
		f[v/k] = 0
	}

	for x := mx; x >= 1; x-- {
		for _, v := range a {
			v /= k
			f[g[x][v]] = min(f[g[x][v]], f[x]+1)
		}
	}

	ans := max(0, f[1]-1)
	for _, v := range a {
		if v > k {
			ans++
		}
	}
	return ans
}

const inf = 1 << 60

func gcd(a, b int) int {
	a, b = max(a, b), min(a, b)
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
