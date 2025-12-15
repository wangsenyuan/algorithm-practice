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
	var t, l, r int
	fmt.Fscan(reader, &t, &l, &r)
	return solve(t, l, r)
}

const mod = 1000000007

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

func solve(t int, l int, r int) int {
	n := r + 1
	var primes []int
	lpf := make([]int, n)
	for i := 2; i < n; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= n {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	f := make([]int, n)
	for i := 2; i < n; i++ {
		if lpf[i] == i {
			// 只能全部人在一组内
			f[i] = i * (i - 1) / 2
		} else {
			x := lpf[i]
			// (x - 1) * i / 2 (x越小越好)
			f[i] = x*(x-1)/2*i/x + f[i/x]
		}
		f[i] %= mod
	}

	ans := f[r]
	for i := r - 1; i >= l; i-- {
		ans = mul(t, ans)
		ans = add(ans, f[i])
	}

	return ans
}
