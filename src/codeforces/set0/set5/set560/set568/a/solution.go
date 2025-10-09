package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var p, q int
	fmt.Fscan(reader, &p, &q)
	res := solve(p, q)
	fmt.Println(res)
}

func solve(p int, q int) int {
	g := gcd(p, q)
	p /= g
	q /= g
	var pi, rub int

	var res int

	lpf := make([]int, 2e6+1)
	var primes []int

	for i := 1; i <= 2e6; i++ {
		if i > 1 && lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
			pi++
		}
		if i > 1 {
			for _, p := range primes {
				if i*p > 2e6 {
					break
				}
				lpf[i*p] = p
				if i%p == 0 {
					break
				}
			}
		}
		if checkPalindrome(i) {
			rub++
		}
		if pi*q <= p*rub {
			res = i
		}
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func checkPalindrome(n int) bool {
	var ds []int

	for i := n; i > 0; i /= 10 {
		ds = append(ds, i%10)
	}
	for l, r := 0, len(ds)-1; l < r; l, r = l+1, r-1 {
		if ds[l] != ds[r] {
			return false
		}
	}
	return true
}
