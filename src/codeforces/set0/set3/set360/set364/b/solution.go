package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0], res[1])
}

func drive(reader *bufio.Reader) []int {
	var n, d int
	fmt.Fscan(reader, &n, &d)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	return solve(d, c)
}

func solve(d int, c []int) []int {
	md := 500000
	dp := make([]int, md+1)
	dp[0] = 1
	for _, v := range c {
		for j := md; j >= v; j-- {
			dp[j] |= dp[j-v]
		}
	}
	var sum int
	var steps int

	for {
		ok := false
		for i := min(sum+d, md); i > sum; i-- {
			if dp[i] != 0 {
				ok = true
				sum = i
				steps++
				break
			}
		}
		if !ok {
			break
		}
	}
	return []int{sum, steps}
}
