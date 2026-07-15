package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		a := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[j])
		}
		res[i] = solve(k, a)
	}
	return res
}

const mod = 998244353

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

func inv(a int) int {
	return pow(a, mod-2)
}

const N = 200010

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	I[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(F[i-1], i)
	}
	I[N-1] = inv(F[N-1])
	for i := N - 2; i > 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}
}

func nCr(n, r int) int {
	if n < r || r < 0 {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(k int, a []int) int {
	n := len(a)

	res := 1
	for i := 0; i < n; {
		x := a[i] % k
		if x == 0 || 2*x == k {
			var arr []int
			for i < n && a[i]%k == x {
				arr = append(arr, a[i])
				i++
			}
			slices.Sort(arr)
			tmp := F[len(arr)]
			var cnt int
			for j := 0; j < len(arr); j++ {
				if j > 0 && arr[j] != arr[j-1] {
					tmp = mul(tmp, I[cnt])
					cnt = 0
				}
				cnt++
			}
			tmp = mul(tmp, I[cnt])
			res = mul(res, tmp)
		} else {
			y := k - x
			cnt := []int{0, 0}
			for i < n && (a[i]%k == x || a[i]%k == y) {
				if a[i]%k == x {
					cnt[0]++
				} else {
					cnt[1]++
				}
				i++
			}
			tmp := nCr(cnt[0]+cnt[1], cnt[0])
			res = mul(res, tmp)
		}
	}

	return res
}
