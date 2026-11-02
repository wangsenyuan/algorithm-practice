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

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int64 {
	// Opposite parities never share a key. Same parity shares a key
	// exactly when the starts differ by 2 or 4.
	groups := [2]map[int][]int{make(map[int][]int), make(map[int][]int)}
	for i := 0; i+4 < len(a); i++ {
		sum := a[i] + a[i+2] - a[i+4]
		groups[i&1][sum] = append(groups[i&1][sum], i)
	}

	sums := map[int]struct{}{}
	for _, group := range groups {
		for sum := range group {
			sums[sum] = struct{}{}
		}
	}

	var ans int64
	for sum := range sums {
		left, right := groups[0][sum], groups[1][sum]
		ans += int64(len(left)) * int64(len(right))
		ans += distantPairs(left)
		ans += distantPairs(right)
	}
	return ans
}

func distantPairs(idx []int) int64 {
	var ans int64
	j := 0
	for i := range idx {
		for j < len(idx) && idx[j] < idx[i]+6 {
			j++
		}
		ans += int64(len(idx) - j)
	}
	return ans
}
