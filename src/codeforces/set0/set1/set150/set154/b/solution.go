package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, line := range drive(reader) {
		fmt.Fprintln(writer, line)
	}
}

func drive(reader *bufio.Reader) []string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	ops := make([][]int, m)
	for i := range m {
		var op string
		var x int
		fmt.Fscan(reader, &op, &x)
		sign := 1
		if op == "-" {
			sign = -1
		}
		ops[i] = []int{sign, x}
	}
	return solve(n, ops)
}

func solve(n int, ops [][]int) []string {
	var primes []int
	lpf := make([]int, n+1)
	ord := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			ord[i] = len(primes)
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i > n {
				break
			}
			lpf[p*i] = p
			if i%p == 0 {
				break
			}
		}
	}

	status := make([]int, len(primes))
	for i := range status {
		status[i] = -1
	}

	ans := make([]string, len(ops))

	flag := make([]int, n+1)

	check := func(x int) int {
		for x > 1 {
			v := lpf[x]
			if status[ord[v]] >= 0 {
				return status[ord[v]]
			}
			x /= v
		}
		return -1
	}

	play := func(x int, on bool) {
		for i := x; i > 1; {
			v := lpf[i]
			if on {
				status[ord[v]] = x
			} else {
				status[ord[v]] = -1
			}
			i /= v
		}
	}

	for i, op := range ops {
		sign, x := op[0], op[1]
		if sign == 1 {
			if flag[x] == 1 {
				ans[i] = "Already on"
			} else {
				j := check(x)
				if j >= 0 {
					ans[i] = fmt.Sprintf("Conflict with %d", j)
				} else {
					play(x, true)
					flag[x] = 1
					ans[i] = "Success"
				}
			}
		} else {
			// sign = -1
			if flag[x] == 0 {
				ans[i] = "Already off"
			} else {
				play(x, false)
				flag[x] = 0
				ans[i] = "Success"
			}
		}
	}

	return ans
}
