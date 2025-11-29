package main

import (
	"bufio"
	"fmt"
	"os"
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
	n := len(a)
	var sum []int
	sum = append(sum, 0)
	for _, v := range a {
		sum = append(sum, sum[len(sum)-1]+v)
	}
	if sum[n]%2 == 1 {
		return false
	}
	half := sum[n] / 2

	pos := make(map[int]int)

	for i, v := range a {
		pos[v] = i
	}

	var pref int
	for i := range n {
		if pref > half {
			break
		}

		if j, ok := pos[half-pref]; ok && i <= j {
			return true
		}
		// half + a[i]

		j := sort.Search(len(sum), func(j int) bool {
			return sum[j] >= half+a[i]
		})

		if j <= n && sum[j] == half+a[i] {
			return true
		}

		pref += a[i]
	}

	return false
}
