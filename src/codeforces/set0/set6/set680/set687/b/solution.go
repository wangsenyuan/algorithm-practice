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
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	return solve(k, c)
}

func solve(k int, c []int) bool {
	x := slices.Max(c)
	x = max(x, k)

	lpf := make([]int, x+1)
	var primes []int
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i > x {
				break
			}
			lpf[p*i] = p

			if i%p == 0 {
				break
			}
		}
	}

	freq := make([]int, len(primes))

	play := func(num int) {
		for num > 1 {
			var w int
			f := lpf[num]
			for num%f == 0 {
				w++
				num /= f
			}
			j := sort.SearchInts(primes, f)
			freq[j] = max(freq[j], w)
		}
	}

	for _, v := range c {
		play(v)
	}

	for k > 1 {
		var w int
		f := lpf[k]
		for k%f == 0 {
			w++
			k /= f
		}
		j := sort.SearchInts(primes, f)
		if w > freq[j] {
			return false
		}
	}

	return true
}
