package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	q := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &q[i])
	}
	return solve(a, q)
}

func solve(a []int, q []int) []int {
	x := slices.Max(a)
	lpf := make([]int, x+1)
	var primes []int
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			primes = append(primes, i)
			lpf[i] = i
		}
		for _, j := range primes {
			if i*j > x {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	cnt := make([]int, x+1)
	freq := make([]int, x+1)

	var sum int
	var tot int

	update := func(v int, d int) {
		var fs []int
		for u := v; u > 1; u /= lpf[u] {
			if freq[lpf[u]] == 0 {
				fs = append(fs, lpf[u])
			}
			freq[lpf[u]]++
		}

		k := len(fs)
		// k <= 7
		for mask := 1; mask < 1<<k; mask++ {
			var parity int
			y := 1
			for i := range k {
				if (mask>>i)&1 == 1 {
					parity ^= 1
					y *= fs[i]
				}
			}

			if d < 0 {
				cnt[y] += d
			}

			if parity == 1 {
				sum += cnt[y] * d
			} else {
				sum -= cnt[y] * d
			}
			if d > 0 {
				cnt[y] += d
			}
		}

		for _, u := range fs {
			freq[u] = 0
		}
	}
	n := len(a)
	seen := make([]bool, n)
	ans := make([]int, len(q))

	for i, j := range q {
		j--
		if seen[j] {
			update(a[j], -1)
			tot--
		} else {
			update(a[j], 1)
			tot++
		}
		seen[j] = !seen[j]
		ans[i] = tot*(tot-1)/2 - sum
	}
	return ans
}
