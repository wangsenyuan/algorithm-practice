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
		_, points, order := drive(reader)
		fmt.Fprintln(writer, points)
		s := fmt.Sprintf("%v", order)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (x int, best int, prices []int) {
	var n int
	fmt.Fscan(reader, &n, &x)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	best, prices = solve(x, a)
	return
}

func solve(x int, a []int) (int, []int) {
	slices.Sort(a)

	var sum, score int
	var prices []int
	n := len(a)

	for i, j := 0, n; i < j; {
		if sum+a[j-1] >= x {
			prices = append(prices, a[j-1])
			score += a[j-1]
			sum += a[j-1]
			j--
		} else {
			prices = append(prices, a[i])
			sum += a[i]
			i++
		}
		sum %= x
	}
	return score, prices
}
