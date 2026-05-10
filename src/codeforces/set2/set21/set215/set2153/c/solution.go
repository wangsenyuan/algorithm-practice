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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	var sum int
	slices.Sort(a)
	n := len(a)
	var cnt int
	var arr []int
	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		sum += (i - j) / 2 * 2 * a[j]
		cnt += (i - j) / 2 * 2
		if (i-j)%2 == 1 {
			arr = append(arr, a[j])
		}
	}

	if cnt == 0 {
		return 0
	}

	slices.Sort(arr)
	slices.Reverse(arr)

	for i := 0; i+1 < len(arr); i++ {
		a, b := arr[i], arr[i+1]
		if sum+b > a {
			return sum + a + b
		}

	}

	for i := 0; i < len(arr); i++ {
		a := arr[i]
		if sum > a {
			return sum + a
		}
	}

	if cnt == 2 {
		return 0
	}

	return sum
}
