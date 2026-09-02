package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int) int {

	w := sort.Search(n, func(i int) bool {
		if i == 0 {
			return false
		}
		if i > n/i || i*i > n {
			return true
		}
		// i * i < n
		if i > n/(i*i) || i*i*i > n {
			return true
		}

		return false
	})
	w--
	if w < 2 {
		return 0
	}

	lpf := make([]int, w+1)
	var primes []int
	for i := 2; i <= w; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p > w {
				break
			}
			lpf[i*p] = p
			if i%p == 0 {
				break
			}
		}
	}

	var ans int
	for i, p := range primes {
		p3 := p * p * p
		l := sort.Search(len(primes), func(j int) bool {
			if primes[j] > n/p3 || primes[j]*p3 > n {
				return true
			}
			return false
		})
		ans += min(l, i)
	}

	return ans
}
