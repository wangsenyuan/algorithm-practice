package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, x)
}

const mod = 1_000_000_007

const N = 500010

var lpf [N]int
var primes []int

func init() {
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= N {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(a []int, x int) int {
	f := make(map[int]int)

	for _, v := range a {
		for v > 1 {
			f[lpf[v]]++
			v /= lpf[v]
		}
	}

	ans := 1

	for _, v := range f {
		ans = mul(ans, v+1)
	}

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
