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

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	// 奇数和偶数配。
	cnt := make([]int, 2)
	for _, v := range a {
		cnt[v&1]++
	}

	if cnt[0] >= cnt[1] {
		return cnt[1]
	}

	res := cnt[0]
	cnt[1] -= cnt[0]

	res += cnt[1] / 3
	return res
}
