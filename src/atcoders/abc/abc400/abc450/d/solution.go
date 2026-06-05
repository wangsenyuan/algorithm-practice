package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	return solve(a, k)
}

func solve(a []int, k int) int {
	n := len(a)
	arr := make([]int, n)
	for i := range n {
		arr[i] = a[i] % k
	}
	slices.Sort(arr)
	best := arr[n-1] - arr[0]

	for range n {
		x := arr[0]
		arr = arr[1:]
		arr = append(arr, x+k)
		best = min(best, arr[n-1]-arr[0])
	}

	return best
}
