package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var l, r, x, y int
	fmt.Fscan(reader, &l, &r, &x, &y)
	return solve(l, r, x, y)
}

func solve(l int, r int, x int, y int) int {
	if x == y {
		if l <= x && x <= r {
			return 1
		}
		return 0
	}
	if y%x != 0 {
		return 0
	}

	// 如果w是一个质数
	w := y / x

	mw := min(w, 1e6)

	lpf := make([]int, mw+1)
	var primes []int
	for i := 2; i <= mw; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > mw {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	var factors []int

	for _, v := range primes {
		if v > w {
			break
		}
		if w%v == 0 {
			v1 := 1
			for w%v == 0 {
				w /= v
				v1 *= v
			}
			factors = append(factors, v1)
		}
	}
	if w > 1 {
		factors = append(factors, w)
	}

	var ans int

	// len(freq) <= 20
	n := len(factors)
	T := 1 << n
	for mask := range T {
		a, b := x, x
		for i := range n {
			if (mask>>i)&1 == 0 {
				a *= factors[i]
			} else {
				b *= factors[i]
			}
		}

		if a >= l && a <= r && b >= l && b <= r {
			ans++
		}
	}

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
