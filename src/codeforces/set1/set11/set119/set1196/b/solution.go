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

	var q int
	fmt.Fscan(reader, &q)
	for range q {
		ends := drive(reader)
		if ends == nil {
			fmt.Fprintln(writer, "NO")
			continue
		}
		fmt.Fprintln(writer, "YES")
		for i, end := range ends {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, end)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) []int {
	n := len(a)
	var res []int
	var sum int
	for i, v := range a {
		if v%2 == 1 {
			res = append(res, i+1)
			sum++
		}
	}
	if len(res) >= k && (sum-k)%2 == 0 {
		res = res[:k-1]
		res = append(res, n)
		return res
	}
	return nil
}
