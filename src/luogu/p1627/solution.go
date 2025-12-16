package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, med int
	fmt.Fscan(reader, &n, &med)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(med, a)
}

func solve(med int, a []int) int {
	n := len(a)
	freq := make([]int, 2*n)
	freq[n]++

	var sum int
	var i int
	for i < n && a[i] != med {
		if a[i] < med {
			sum--
		} else {
			sum++
		}
		freq[n+sum]++
		i++
	}
	var res int
	for i < n {
		if a[i] < med {
			sum--
		} else if a[i] > med {
			sum++
		}
		res += freq[n+sum]
		i++
	}

	return res
}
