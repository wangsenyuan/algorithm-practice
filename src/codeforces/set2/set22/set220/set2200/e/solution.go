package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

const X = 1e6 + 10

var lpf [X]int
var primeFactorCount [X]int

func init() {
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			for j := i; j < X; j += i {
				lpf[j] = i
				primeFactorCount[j]++
			}
		}
	}
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) string {
	if sort.IntsAreSorted(a) {
		return "Bob"
	}

	var mxPrimeFactor int
	for _, num := range a {
		if num == 1 {
			// alice搞出一个更大的质数，肯定能获胜，最后的结果肯定不会有序
			if mxPrimeFactor > 1 {
				return "Alice"
			}
			continue
		}
		if primeFactorCount[num] > 1 {
			return "Alice"
		}

		if mxPrimeFactor > lpf[num] {
			return "Alice"
		}
		mxPrimeFactor = max(mxPrimeFactor, lpf[num])
	}
	return "Bob"
}
