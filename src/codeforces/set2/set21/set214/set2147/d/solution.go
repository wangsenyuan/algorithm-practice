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
		fmt.Fprintf(writer, "%d %d\n", res[0], res[1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	var sum int
	slices.Sort(a)
	var odd []int
	for i := 0; i < len(a); {
		j := i
		for i < len(a) && a[i] == a[j] {
			i++
		}
		sum += (i - j) * a[j]
		if a[j]%2 == 1 {
			odd = append(odd, i-j)
		}
	}
	slices.Sort(odd)
	slices.Reverse(odd)

	var diff int
	for i, v := range odd {
		if i&1 == 0 {
			diff += v
		} else {
			diff -= v
		}
	}
	// a + b = sum
	// a - b = diff
	alice := (sum + diff) / 2
	return []int{alice, sum - alice}
}
