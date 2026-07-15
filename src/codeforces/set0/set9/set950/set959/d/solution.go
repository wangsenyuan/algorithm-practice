package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := drive(reader)
	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	const X = 2000000
	var lpf [X]int
	var primes []int
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= X {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	vis := make([]bool, X)
	n := len(a)
	b := make([]int, n)

	check := func(v int) bool {
		for v > 1 {
			w := lpf[v]
			if vis[w] {
				return false
			}
			for v%w == 0 {
				v /= w
			}
		}
		return true
	}

	for i := range n {
		// a[i] > 1 holds
		if check(a[i]) {
			for v := a[i]; v > 1; v /= lpf[v] {
				vis[lpf[v]] = true
			}
			b[i] = a[i]
			continue
		}
		// b[i] 不一定是一个素数, 它只需要是 a[i]+1开始处理就可以了

		x := a[i] + 1
		for !check(x) {
			x++
		}
		b[i] = x

		for v := x; v > 1; v /= lpf[v] {
			vis[lpf[v]] = true
		}

		for j := i + 1; j < n; j++ {
			for len(primes) > 0 && vis[primes[0]] {
				primes = primes[1:]
			}
			b[j] = primes[0]
			vis[primes[0]] = true
			primes = primes[1:]
		}
		break
	}

	return b
}
