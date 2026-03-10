package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var k int
	fmt.Fscan(reader, &k)
	a := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	x := slices.Max(a)
	lpf := make([]int, x+1)
	var primes []int
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			primes = append(primes, i)
			lpf[i] = i
		}
		for _, p := range primes {
			if i*p > x {
				break
			}
			lpf[i*p] = p
			if i%p == 0 {
				break
			}
		}
	}

	cnt := make([]int, x+1)
	for _, v := range a {
		cnt[v]++
	}
	for i := x - 1; i > 1; i-- {
		cnt[i] += cnt[i+1]
	}

	for i := x; i > 1; i-- {
		if lpf[i] != i {
			cnt[lpf[i]] += cnt[i]
		}
		cnt[i/lpf[i]] += cnt[i]
	}

	check := func(n int) bool {
		if n < x {
			return false
		}
		for _, p := range primes {
			w := p
			var cur int
			for w <= n {
				cur += n / w
				if w > n/p || w*p > n {
					break
				}
				w *= p
			}
			if cur < cnt[p] {
				return false
			}
		}
		return true
	}
	var sum int
	for _, v := range a {
		sum += v
	}

	return sort.Search(sum, check)
}
