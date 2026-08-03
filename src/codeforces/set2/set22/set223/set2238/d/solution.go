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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

const N = 1e6 + 10

var lpf [N]int

func init() {
	var primes []int
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= N {
				break
			}
			lpf[p*i] = p
			if i%p == 0 {
				break
			}
		}
	}
}

func solve(n int) int {
	var sum int
	var cnt int
	for n > 1 {
		x := lpf[n]
		cnt++
		for n%x == 0 {
			sum++
			n /= x
		}
	}

	return sum + cnt - 1
}
