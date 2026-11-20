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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		colors := drive(reader)
		m := 0
		for _, c := range colors {
			m = max(m, c)
		}
		fmt.Fprintln(writer, m)
		for i, c := range colors {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, c)
		}
		fmt.Fprintln(writer)
	}
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
	n := len(a)
	c := slices.Max(a)
	color := make([]int, c+1)
	ans := make([]int, n)

	var primes []int
	lpf := make([]int, c+1)
	for i := 2; i <= c; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > c {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	var cnt int
	for i, v := range a {
		d := lpf[v]
		if color[d] != 0 {
			ans[i] = color[d]
			continue
		}
		cnt++
		color[d] = cnt
		ans[i] = cnt
	}
	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
