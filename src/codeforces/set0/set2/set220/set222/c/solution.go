package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintln(writer, len(res[0]), len(res[1]))
	for _, v := range res[0] {
		fmt.Fprint(writer, v, " ")
	}
	fmt.Fprintln(writer)
	for _, v := range res[1] {
		fmt.Fprint(writer, v, " ")
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) [][]int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) [][]int {
	x := max(slices.Max(a), slices.Max(b))

	lpf := make([]int, x+1)
	var primes []int
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
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

	// m := len(primes)
	freq := make([]int, x+1)
	for _, num := range a {
		for num > 1 {
			freq[lpf[num]]++
			num /= lpf[num]
		}
	}
	for _, num := range b {
		for num > 1 {
			freq[lpf[num]]--
			num /= lpf[num]
		}
	}

	for i, v := range a {
		u := v
		for v > 1 {
			x := lpf[v]
			if freq[x] > 0 {
				// keep it
				freq[x]--
			} else {
				u /= x
			}
			v /= x
		}
		a[i] = u
	}

	for i, v := range b {
		u := v
		for v > 1 {
			x := lpf[v]
			if freq[x] < 0 {
				freq[x]++
			} else {
				u /= x
			}
			v /= x
		}
		b[i] = u
	}

	return [][]int{a, b}
}
