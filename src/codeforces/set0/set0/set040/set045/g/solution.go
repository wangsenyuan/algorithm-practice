package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	res := solve(n)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int) []int {
	if n < 2 {
		return nil
	}
	lpf := make([]int, n+1)
	var primes []int
	for i := 2; i <= n; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > n {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	sum := n * (n + 1) / 2

	if checkPrimeFast(sum) {
		ans := make([]int, n)
		for i := range n {
			ans[i] = 1
		}
		return ans
	}

	playEven := func() []int {
		var s1 int
		for _, p := range primes {
			if checkPrimeFast(sum - p) {
				s1 = p
				break
			}
		}
		if s1 == 0 {
			for i := n + 1; i < sum; i++ {
				if checkPrimeFast(i) && checkPrimeFast(sum-i) {
					s1 = i
					break
				}
			}
		}
		if s1 == 0 {
			return nil
		}

		ans := make([]int, n)
		for i := range n {
			ans[i] = 1
		}

		for i := n; i > 0 && s1 > 0; i-- {
			if s1 >= i {
				ans[i-1] = 2
				s1 -= i
			}
		}

		return ans
	}

	if sum&1 == 0 {
		return playEven()
	}

	// sum is odd
	if checkPrimeFast(sum - 2) {
		ans := make([]int, n)
		for i := range n {
			ans[i] = 1
		}
		ans[1] = 2
		return ans
	}
	if n < 4 {
		return nil
	}
	ans := make([]int, n)
	ans[0] = 1
	ans[1] = 1
	ans[3] = 1
	sum -= 7
	ok := false
	for _, p := range primes {
		if p == 2 {
			continue
		}
		if checkPrimeFast(sum - p) {
			ans[p-1] = 2
			ok = true
			break
		}
	}

	if !ok {
		return nil
	}
	for i := range n {
		if ans[i] == 0 {
			ans[i] = 3
		}
	}

	return ans
}

func mul(a, b, mod int) int {
	return a * b % mod
}

func pow(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a, mod)
		}
		a = mul(a, a, mod)
		b >>= 1
	}
	return res
}

func checkPrimeFast(n int) bool {
	if n < 2 {
		return false
	}
	for _, p := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37} {
		if n%p == 0 {
			return n == p
		}
	}
	d := n - 1
	var s int
	for d&1 == 0 {
		d >>= 1
		s++
	}
	for _, a := range []int{2, 325, 9375, 28178, 450775, 9780504, 1795265022} {
		if a%n == 0 {
			continue
		}
		x := pow(a%n, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		comp := true
		for r := 1; r < s; r++ {
			x = mul(x, x, n)
			if x == n-1 {
				comp = false
				break
			}
		}
		if comp {
			return false
		}
	}
	return true
}
