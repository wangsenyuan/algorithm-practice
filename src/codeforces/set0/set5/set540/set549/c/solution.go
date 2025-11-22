package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, k, a)
}

func solve(n int, k int, a []int) string {
	cnt := make([]int, 2)
	for _, v := range a {
		cnt[v&1]++
	}

	if n == k {
		if cnt[1]&1 == 0 {
			return "Daenerys"
		}
		return "Stannis"
	}

	if cnt[1] <= (n-k)/2 {
		return "Daenerys"
	}

	if (n-k)&1 == 1 {
		if k&1 == 0 {
			if cnt[0] <= (n-k)/2 {
				return "Daenerys"
			}
		}
		return "Stannis"
	}

	if k&1 == 0 {
		return "Daenerys"
	}

	if cnt[0] <= (n-k)/2 {
		return "Stannis"
	}

	return "Daenerys"
}
