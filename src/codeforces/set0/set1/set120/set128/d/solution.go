package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) bool {
	arr := slices.Clone(a)
	sort.Ints(arr)
	arr = slices.Compact(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1]+1 {
			return false
		}
	}
	m := len(arr)
	if m == 1 {
		return len(a) == 1
	}
	x := arr[0]
	freq := make([]int, m)
	for _, v := range a {
		freq[v-x]++
	}

	for i := m - 2; i > 0; i-- {
		if freq[i+1] > freq[i] {
			return false
		}
		freq[i] -= freq[i+1]
		if freq[i] == 0 {
			return false
		}
	}

	return freq[0] == freq[1]
}
