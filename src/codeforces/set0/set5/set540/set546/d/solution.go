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
	var t int
	fmt.Fscan(reader, &t)
	for range t {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		fmt.Fprintln(writer, solve(a, b))
	}
}

const X = 5_000_010

var divs_count [X]int
var divs_sum [X]int

func init() {
	var primes []int
	lpf := make([]int, X)
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= X {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
	for i := 2; i < X; i++ {
		j := lpf[i]
		divs_count[i] = divs_count[i/j] + 1
		divs_sum[i] = divs_sum[i-1] + divs_count[i]
	}
}

func solve(a int, b int) int {
	return divs_sum[a] - divs_sum[b]
}
