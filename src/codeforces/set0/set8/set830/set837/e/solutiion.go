package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscan(reader, &a, &b)
	fmt.Println(solve(a, b))
}

func solve(a int, b int) int {

	n := min(a, 2000000)

	var primes []int
	set := make([]bool, n+1)
	for i := 2; i <= n; i++ {

		if !set[i] {
			primes = append(primes, i)
			set[i] = true
		}

		for _, j := range primes {
			if j*i > n {
				break
			}
			set[i*j] = true
			if i%j == 0 {
				break
			}
		}
	}

	a1 := a
	for _, v := range primes {
		if a1 == 1 {
			break
		}
		for a1%v == 0 {
			a1 /= v
		}
	}
	if a1 > 1 {
		primes = append(primes, a1)
	}

	checkPrime := func(x int) bool {
		i := sort.SearchInts(primes, x)
		return i < len(primes) && primes[i] == x
	}

	var f func(a, b int) int

	f = func(a, b int) int {
		if b == 0 {
			return 0
		}
		if a == 1 || checkPrime(a) {
			return b%a + b/a
		}

		var res int

		for gcd(a, b) == 1 && b > 0 {
			res++
			b--
		}

		c := gcd(a, b)
		a /= c
		b /= c

		return res + f(a, b)
	}

	return f(a, b)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
