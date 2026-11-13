package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	if len(a) == 1 {
		return a[0] % mod
	}
	// 奇数长度先做一行，下一行长度是偶数，并且以加法开头
	if len(a)%2 == 1 {
		b := make([]int, len(a)-1)
		sign := 1
		for i := range b {
			b[i] = add(a[i], mul(sign, a[i+1]))
			sign = -sign
		}
		a = b
	}

	var first []int
	var second []int
	for i, v := range a {
		if i&1 == 0 {
			first = append(first, v)
		} else {
			second = append(second, v)
		}
	}

	x := play(first)
	y := play(second)
	// n%4 == 0 时最后一步是减法，n%4 == 2 时是加法
	if len(a)%4 == 0 {
		return sub(x, y)
	}
	return add(x, y)
}

func play(nums []int) int {
	n := len(nums)
	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(F[i-1], i)
	}
	I := make([]int, n+1)
	I[n] = pow(F[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}

	nCr := func(n int, r int) int {
		return mul(F[n], I[r], I[n-r])
	}

	var ans int
	for i := range n {
		ans = add(ans, mul(nCr(n-1, i), nums[i]))
	}
	return ans

}

const mod = 1e9 + 7

func add(nums ...int) int {
	var res int
	for _, num := range nums {
		res = (res + num) % mod
		if res < 0 {
			res += mod
		}
	}
	return res
}

func sub(a int, b int) int {
	return add(a, mod-b)
}

func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res = (res * num) % mod
		if res < 0 {
			res += mod
		}
	}
	return res
}

func pow(a int, b int) int {
	res := 1
	for b > 0 {
		if b&1 > 0 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}
